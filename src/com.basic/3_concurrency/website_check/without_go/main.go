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

	for _, website := range listOfWebsites {
		checkWebsite(website)
	}

}

func checkWebsite(website string) {
	_, err := http.Get(website)
	if err != nil {
		fmt.Println(website, "might be down!")
		return
	}
	fmt.Println(website, "is up!")
}
