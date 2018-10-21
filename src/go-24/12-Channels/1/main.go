package main

import (
	"fmt"
	"time"
)

func ping1(ch chan string) {
	for {
		time.Sleep(time.Second * 3)
		ch <- "Ping from channel 1"
	}
}

func ping2(ch chan string) {
	for {
		time.Sleep(time.Second * 2)
		ch <- "Ping from channel 2"
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go ping1(ch1)
	go ping2(ch2)

	for {

		select {
		case msg1 := <-ch1:
			fmt.Println("recieved", msg1)
		case msg2 := <-ch2:
			fmt.Println("recieved", msg2)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("No messages recieved. giving up.")
		}
	}

}
