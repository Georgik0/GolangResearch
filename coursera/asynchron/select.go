package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)
	//ch2 <- 2
	ch1 <- 1
	ch1 <- 2
	select {
	case r := <-ch1:
		fmt.Println("read from ch1:", r)
	case ch2 <- 1:
		fmt.Println("Write in ch2: 1")
	default:
		fmt.Println("Default")
	}
}
