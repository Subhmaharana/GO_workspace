package customapi

import (
	"encoding/json"
	"net/http"
)

//Book type with Name, Author and ISBN
type Book struct {
	Title  string
	Author string
}

var Books = []Book{
	Book{
		Title:  "oolala",
		Author: "ccimwewcm",
	},
	Book{
		Title:  "oolala2",
		Author: "ccimwewc2m",
	},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
