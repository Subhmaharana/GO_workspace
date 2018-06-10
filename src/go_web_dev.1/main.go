package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello web dev")
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
