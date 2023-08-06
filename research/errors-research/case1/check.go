package case1

import (
	"errors"
	"fmt"
)

const errText = "some text"

type myErrType struct {
	msg string
}

func (m myErrType) Error() string {
	return m.Error()
}

func Check() {
	IsCompare()
}

func AsCompare() {
	myErr := &myErrType{errText}
	myErr2 := &myErrType{errText + "1"}
	err := errors.New(errText)
	err2 := errors.New(errText + "1")

	if errors.As(err, myErr) {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR compare")
	}

	if errors.As(*myErr, myErr2) {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR compare")
	}

	if errors.As(err, &err2) {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR compare")
	}

	if err == *myErr {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR compare")
	}
}

func IsCompare() {
	myErr := &myErrType{errText}
	myErr2 := &myErrType{errText}
	err := errors.New(errText)
	err2 := errors.New(errText)

	if errors.Is(err, myErr) {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR compare")
	}

	if errors.Is(myErr, myErr2) {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR compare")
	}

	if errors.Is(err, err2) {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR compare")
	}

	if err == *myErr {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR compare")
	}
}
