package in

import "fmt"

type in struct {
	a int
}

func GetIn() *in {
	return &in{}
}

func (i *in) Check() {
	fmt.Println("Check")
}
