package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum //send sum to c
}

func main() {
	a := []int{1, 1, 1, 1, 2, 2, 2, 2}

	c := make(chan int)

	//here we are asynchrounously calling the sum function by splitting the array into two
	go sum(a[:len(a)/2], c) //first part
	go sum(a[len(a)/2:], c) //second part

	x, y := <-c, <-c

	fmt.Println(x, "+", y, "=", x+y)

}
