package main

import (
	"fmt"
)

func main() {
	// TRY: looping 1st

	// make channel

	// send data to channels
	// close channel

	// loop over range of channel elements
	// print elements

	count := 5
	channel := make(chan int, 5)

	for i := 0; i < count; i++ {
		channel <- i
	}
	close(channel)

	for i := range channel {
		fmt.Println(i)
	}
}
