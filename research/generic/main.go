package main

import (
	"fmt"
)

type (
	type1 int64
	type2 map[string]string
	type3 map[type1]type2

	someType struct {
		t1 type1
		t2 type2
		t3 type3
	}
)

func main() {
	s := someType{t2: type2{
		"1": "1",
		"2": "2",
	}}

	newS := transform[type2](s)
	fmt.Println(newS["1"])
	fmt.Println(interface{}(newS).(type2))
}

func transform[T any](s someType) T {
	return interface{}(s.t2).(T)
}

func funcInInterface[T any](f T) {
	var zero interface{}
	var someFunc func()

	switch someFunc.(type) {

	}
}
