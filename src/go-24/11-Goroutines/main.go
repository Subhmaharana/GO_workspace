package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	urls := make([]string, 3)

	urls[0] = "https://www.amazon.in"
	urls[1] = "https://www.amazon.com"
	urls[2] = "https://www.usa.gov"

	for _, u := range urls {
		//the response output will come in order of response time after adding "go"
		go responseTime(u)
	}

	time.Sleep(time.Second * 5)

}

func responseTime(url string) {
	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	elasped := time.Since(start).Seconds()

	fmt.Printf("%s took %v seconds\n", url, elasped)
}
