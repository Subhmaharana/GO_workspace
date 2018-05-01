package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	listOfWebsites := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
		"localhost:4200",
	}

	//channel of type string
	ch := make(chan string)

	for _, website := range listOfWebsites {
		go checkWebsite(website, ch)
	}

	//alternate of infinite for loop(looping through the channel)
	for l := range ch {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkWebsite(link, ch)
		}(l)
	}
}

func checkWebsite(website string, ch chan string) {
	_, err := http.Get(website)
	if err != nil {
		fmt.Println(website, "might be down!")
		ch <- website
		return
	}
	fmt.Println(website, "is up!")
	ch <- website
}
