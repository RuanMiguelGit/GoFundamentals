package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
}

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 // This line is causing the issue
	ch <- 5 // This line is causing the issue
	fmt.Println("Executou")
	ch <- 6 // This line is causing the issue
	ch <- 7 // This line is causing the issue
}
