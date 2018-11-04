package main

//Book type with Name, Author and ISBN
type Book struct {
	Title  string
	Author string
}

var books = map[string]Book{
	"123456789": Book{
		Title:  "oolala",
		Author: "ccimwewcm",
	},
	"987654321": Book{
		Title:  "oolala2",
		Author: "ccimwewc2m",
	},
}

//AllBooks returns a slice of all books
func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}
