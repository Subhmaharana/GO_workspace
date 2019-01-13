package main

import (
	"fmt"
	"sort"
)

func main() {

	A := []int{4, 12, 15, 19, 20}
	fmt.Println(subsets(A))
}

func subsets(A []int) [][]int {
	sort.Ints(A)

	result := [][]int{{}}
	for i := 0; i < len(A); i++ {
		generate(A, i, []int{}, &result)
	}

	return result
}

func generate(A []int, pos int, p []int, result *[][]int) {
	if pos == len(A) {
		return
	}

	p = append(p, A[pos])
	*result = append(*result, p)

	for i := pos + 1; i < len(A); i++ {
		generate(A, i, p, result)
	}

	//removing last element
	p = p[:len(p)-1]
}
