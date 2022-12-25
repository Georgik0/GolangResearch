package main

import (
	"fmt"
	"time"
)

type data struct {
	end time.Duration
}

func main() {
	d := f()
	fmt.Println("d := f()  -->  d.end = ", d.end)
}

func f() data {
	d := data{}

	start := time.Now()
	defer func() {
		d.end = time.Now().Sub(start)
		fmt.Println("in f() d.end = ", d.end)
	}()

	time.Sleep(2)
	fmt.Println("out of f()")

	return d
}
