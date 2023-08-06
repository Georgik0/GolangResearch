package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	err := test1()

	fmt.Println(err)
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
