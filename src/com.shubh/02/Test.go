package main

import "fmt"
import "sort"

type book float64

func (this book) printTitle() book{
	return this
}

func main() {
	var b book = "Subhransu"
	b.printTitle()
}