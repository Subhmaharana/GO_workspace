package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	fmt.Println("check")
	result := doFibonacci(34)
	fmt.Println("The Result is..", result)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}

func doFibonacci(num int) int {
	if (num < 2) {
		return num
	}
	return doFibonacci(num - 1) + doFibonacci(num - 2)
}


