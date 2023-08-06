package main

import "fmt"

var _ checkerB = (*B)(nil)

func main() {
	var iA checkerA

	modifA := iA.(interface{}).(A)

	modifB := iA.(interface{}).(B)

	fmt.Println(modifA)
	fmt.Println(modifB)
}

type checkerB interface {
	f2()
}

type checkerA interface {
	f1()
}

type A struct {
}

func (a A) f1() {

}

type B struct {
}
