package main

import (
	"fmt"
	"go_simple-microservice/customapi"
	"net/http"
	"os"
)

func main() {

	//here I am handling different type of requests which can come
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo) //api/echo?message=Something
	http.HandleFunc("/api/books", customapi.BooksHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Subhransu...")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
