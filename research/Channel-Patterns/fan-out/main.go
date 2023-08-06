package main

import (
	"math/rand"
	"time"
)

func main() {
	fanOut()
}

func fanOut() {
	children := 2000
	ch := make(chan string, children)

	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		println(d)
	}
	time.Sleep(time.Second)
}
