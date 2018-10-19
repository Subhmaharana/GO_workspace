package main

import (
	"fmt"
)

func main() {
	var cheeses [2]string
	cheeses[0] = "OOlala"

	for _, cheese := range cheeses {
		fmt.Println(cheese)
	}
	fmt.Println("----------------")

	// using slices
	var cheesesSlices = make([]string, 2)
	cheesesSlices[0] = "OOlala"
	cheesesSlices = append(cheesesSlices, "Subh", "Mahara")
	for _, cheese := range cheesesSlices {
		fmt.Println(cheese)
	}

}
