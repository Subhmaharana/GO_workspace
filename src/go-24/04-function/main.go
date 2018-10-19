package main

import (
	"fmt"
)

func main() {
	// recursive function
	eating(0, 0)

	//multiple args returns
	i, j, k := multipleArgFunc(1, 2)
	fmt.Println("Got the values:", i, j, k)
}
func eating(food int, eaten int) int {
	if eaten > 10 {
		fmt.Println("I am Full...:|")
		return eaten
	}
	food = food + 1
	fmt.Println("I am eating..", food)
	eaten = food
	return eating(food, eaten)
}

func multipleArgFunc(a int, b int) (int, int, int) {
	return a, b, a + b
}
