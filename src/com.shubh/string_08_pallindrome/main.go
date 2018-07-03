package main

import (
	"fmt"
)

func main() {
	result := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(result)
}

func isPalindrome(A string) int {
	ret := 1
	s, e := 0, len(A)-1
	for s < e {
		for s < len(A) && !isAlphabetOrDigit(A[s]) {
			s++
		}
		for e >= 0 && !isAlphabetOrDigit(A[e]) {
			e--
		}

		if s > e {
			break
		}

		if toLower(A[s]) != toLower(A[e]) {
			ret = 0
			break
		}
		s++
		e--
	}
	return ret
}

func isAlphabetOrDigit(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
