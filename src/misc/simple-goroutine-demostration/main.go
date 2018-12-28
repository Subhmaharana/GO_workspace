package main

import (
	"fmt"
	"time"
)

func main() {
	go conductSportsEvent("Kho Kho", 10)
	go conductSportsEvent("Kabbadi", 6)
	go conductSportsEvent("BasketBall", 3)

	fmt.Println("End of main function...")

	//Adding a sleep here to prevent exit of the program...
	time.Sleep(time.Duration(10000) * time.Minute)
}

func conductSportsEvent(sport string, timeInSeconds int64) {
	fmt.Println(sport, "started...: Should take ", timeInSeconds)
	time.Sleep(time.Duration(timeInSeconds) * time.Second)
	fmt.Println("Finished ", sport)
}
