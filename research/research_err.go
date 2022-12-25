package main

import (
	"errors"
	"fmt"
)

func CreateRule() error {
	err := validationRule()
	if err != nil {
		return err
	}
	return nil
}

func validationRule() error {
	return errors.New("smthn")
}

func main() {
	var wantErr error
	var err error

	wantErr = errors.New("smthn")
	err = CreateRule()

	fmt.Println("err    :", err)
	fmt.Println("wanrErr:", wantErr)

	fmt.Println(errors.Is(wantErr, err))
	fmt.Println(errors.Is(err, wantErr))
}
