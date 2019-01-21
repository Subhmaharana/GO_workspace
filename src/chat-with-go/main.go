package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", goToHomePage)
	http.ListenAndServe(":8080", nil)
}

func goToHomePage(responseWriter http.ResponseWriter, req *http.Request) {

	fmt.Println("Going to HomePage...")
	http.ServeFile(responseWriter, req, "chat_box.html")

}
