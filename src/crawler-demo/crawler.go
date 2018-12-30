package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

//Adding a map to avoid looping into same website again
var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify a start page..")
		os.Exit(1)
	}

	linkChannel := make(chan string)

	go func() {
		linkChannel <- args[0]
	}()

	for uri := range linkChannel {
		addToQueue(uri, linkChannel)
	}

}

func addToQueue(uri string, queue chan string) {
	visited[uri] = true
	resp, err := http.Get(uri)

	if err != nil {
		fmt.Println("Unable to connect to website..")
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	links := getAllLinks(resp.Body)

	for _, link := range links {
		absoluteURL := fixUrl(link, uri)
		if uri != "" && !visited[absoluteURL] {
			go func() { queue <- absoluteURL }()
		}
	}

}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	fmt.Println(uri.String())
	return uri.String()
}

//Collect all links from response body and return it as an array of strings
func getAllLinks(body io.Reader) []string {
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			//todo: links list shoudn't contain duplicates
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
