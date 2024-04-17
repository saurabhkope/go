package main

import (
	"fmt"
)

func main() {
	// Creating a channel of type int
	ch := make(chan int)

	// Producer: Send some integers to the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// Close the channel after sending all values
		close(ch)
	}()

	// Consumer: Receive values from the channel using a for range loop
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("Channel closed")
}
