package main

import (
	"fmt"
	"time"
)

func main() {
	allEventsDone := make(chan int)

	go conductSportsEvent("Kho Kho", 10, allEventsDone)
	go conductSportsEvent("Kabbadi", 6, allEventsDone)
	go conductSportsEvent("BasketBall", 3, allEventsDone)

	fmt.Println("End of main function...")

	for i := 1; i <= 3; i++ {
		fmt.Println(<-allEventsDone)
	}

	//Adding a sleep here to prevent exit of the program...
	// time.Sleep(time.Duration(10000) * time.Minute)
}

func conductSportsEvent(sport string, timeInSeconds int64, done chan int) {
	fmt.Println(sport, "started...: Should take ", timeInSeconds, " seconds")
	time.Sleep(time.Duration(timeInSeconds) * time.Second)
	fmt.Println("Finished ", sport)
	done <- int(timeInSeconds)
}
