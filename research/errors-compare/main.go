package main

import (
	"errors"
	"fmt"
)

func main() {
	const txt = "some text"

	err1 := errors.New(txt)
	err2 := errors.New(txt)

	fmt.Println(err1 == err2)
	fmt.Println(errors.Is(err1, err2))
	fmt.Println(err1 == err1)
}
