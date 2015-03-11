package main

import (
	"fmt"
)

func main() {
	// done channel has NO buffer / we did not specify its capacity
	// in other words: this channel is synchronous / blocking
	done := make(chan bool)

	go func() {
		fmt.Println("called?")

		done <- true
	}()

	fmt.Println("def called")

	<-done
}
