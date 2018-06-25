package main

import (
	"fmt"
)

func main() {
	result := lengthOfLastWord("Hello World   ")
	fmt.Println(result)
}

/**
 * @input A : String
 *
 * @Output Integer
 */
func lengthOfLastWord(A string) int {
	n := len(A)
	i := n - 1
	WL := 0
	for i >= 0 {
		if A[i] == ' ' && WL == 0 {
			i--
			continue
		} else if A[i] == ' ' {
			return WL
		} else {
			WL++
		}
		i--
	}
	return WL
}
