package main

import "sort"
import "fmt"

func main() {
	intArr := []int{3, 2, 1, 4, 5, 6, 7}
	fmt.Println(wave(intArr))
}

func wave(A []int) ([]int) {
	if len(A) == 1 {
		return A
	}
	sort.Ints(A)

	for i := 1; i < len(A); i += 2 {
		A[i-1], A[i] = A[i], A[i-1]
	}
	return A
}
