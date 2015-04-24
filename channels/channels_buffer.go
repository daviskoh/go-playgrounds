package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	 * TRY creating buffered channel
	 * pass capacity (2nd arg) to make func
	 *
	 * all writing operations are performed without waiting for the first read for the buffer of the channel allows to store all three messages
	 */
	message := make(chan string)
	// message := make(chan string, 3)
	count := 3

	go func() {
		for i := 1; i <= count; i++ {
			fmt.Println("send message")
			message <- fmt.Sprintf("message %d", i)
		}
	}()

	time.Sleep(time.Second * 3)

	for i := 1; i <= count; i++ {
		fmt.Println(<-message)
	}
}
