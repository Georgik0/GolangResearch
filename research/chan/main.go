package main

import (
	"errors"
	"fmt"
	"time"

	"GolangResearch/research/chan/reading"
)

func main() {
	//init_chan.Check()
	reading.Check()
}

func test1() error {
	ch := make(chan error)

	go func() {
		time.Sleep(5)
		ch <- errors.New("err")
	}()

	err := <-ch
	fmt.Println(err)

	return <-ch
}
