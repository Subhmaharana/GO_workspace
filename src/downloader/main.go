package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type download struct {
	Title    string `json:"title"`
	Location string `json:"location"`
}

func status(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("hello status..")
}

func handleDownloadRequest(response http.ResponseWriter, request *http.Request) {
	var downloadRequest download

	r, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, "bad request", 400)
		log.Println(err)
		return
	}
	defer request.Body.Close()

	err = json.Unmarshal(r, &downloadRequest)
	if err != nil {
		http.Error(response, "bad request: "+err.Error(), 400)
		return
	}

	err = getFile(downloadRequest)
	if err != nil {
		http.Error(response, "internal server error", 500)
		return
	}

	log.Printf("%#v", downloadRequest)
}

func getFile(downloadRequest download) error {
	parsedURL, err := url.Parse(downloadRequest.Location)

	if err != nil {
		log.Println(err)
		return err
	}

	response, err := http.Get(downloadRequest.Location)
	if err != nil {
		log.Println(err)
		return err
	}

	defer response.Body.Close()

	out, err := os.Create(filepath.Base(parsedURL.Path))
	_, err = io.Copy(out, response.Body)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func main() {
	fmt.Println("Downloader started...")

	download := download{
		Title:    "title-test",
		Location: "location-test",
	}
	fmt.Printf("%v", download)

	http.HandleFunc("/", status)
	http.HandleFunc("/download", handleDownloadRequest)
	http.ListenAndServe(":3000", nil)
}
