package main

import (
	"math/rand"
	"time"
)

func main() {
	waitForTask()
}

func waitForTask() {
	ch := make(chan string)

	go func() {
		d := <-ch
		println(d)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "data"

	time.Sleep(time.Second)
}
