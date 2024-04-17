package main

import "fmt"

// Function to generate Fibonacci numbers and send them to a channel
func fibonacci(n int, ch chan<- int) {
    a, b := 0, 1
    for i := 0; i < n; i++ {
        ch <- a
        a, b = b, a+b
    }
    close(ch)
}

func main() {
    // Create a channel
    ch := make(chan int)

    // Start a goroutine to generate Fibonacci numbers
    go fibonacci(10, ch)

    // Read Fibonacci numbers from the channel and print them
    for num := range ch {
        fmt.Println(num)
    }
}
