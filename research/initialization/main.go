package main

import (
	"fmt"
	"reflect"
)

func main() {
	case3()
}

func case1() {
	a := [...]string{"no error", "Eio", "invalid argument"}
	b := []string{"no error", "Eio", "invalid argument"}

	info(a)
	info(b)
}

func case2() {
	var aPtr [1]string

	info(aPtr)
}

func case3() {
	bNIL := new([]string)
	var bNIL_var *[]string

	bNILV := reflect.ValueOf(bNIL)
	bNIL_varV := reflect.ValueOf(bNIL_var)
	info(bNILV)
	info(bNIL_varV) // panic

	info(bNIL)
	info(bNIL_var) // panic
}

func info(a interface{}) {
	if aSl, ok := a.([]string); ok {
		fmt.Print(len(aSl), cap(aSl), " ")
	}

	if aSl, ok := a.(*[]string); ok {
		if aSl != nil {
			fmt.Print(len(*aSl), cap(*aSl), " ")
		}
	}
	fmt.Printf("%#v, %#v\n", &a, a)

	fmt.Printf("Type %T\n\n", a)
}
