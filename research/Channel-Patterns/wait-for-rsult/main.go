package main

import (
	"math/rand"
	"time"
)

func main() {
	waitForResult()
}

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
	}()

	d := <-ch

	time.Sleep(time.Second)
	println(d)
}
