package main

import (
	"fmt"
	"net/http"
)

func main() {
	listOfWebsites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	//channel of type string
	ch := make(chan string)

	for _, website := range listOfWebsites {
		go checkWebsite(website, ch)
	}
	fmt.Println(<-ch)

}

func checkWebsite(website string, ch chan string) {
	_, err := http.Get(website)
	if err != nil {
		fmt.Println(website, "might be down!")
		ch <- "might be down i think"
		return
	}
	fmt.Println(website, "is up!")
	ch <- "Yep it is up!"
}
