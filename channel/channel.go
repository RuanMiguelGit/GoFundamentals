package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // writing data to the channel
	fmt.Println("Data written to channel")
	ch <- 2
	fmt.Println(<-ch) // reading data from the channel
}
