package main

import (
	"fmt"
	"reflect"
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
	//transformGl()
	researchImplement()
}

func transformGl() {
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
	// ERROR compile

	//var zero interface{}
	//var someFunc func()
	//
	//switch someFunc.(type) {
	//default:
	//	return
	//}
}

type checker struct {
	a int
}

func (c *checker) write() {
	fmt.Println(c.a)
}

type checkerI interface {
	write()
}

func researchImplement() {
	a1 := int32(5)
	implementNilType(a1)

	a2 := int32(5)
	implementNilType(&a2)

	i := interface{}(&a2)
	implementNilType(i)

	c := checkerI(&checker{a: 5})
	implementNilType(c)
}

func implementNilType[T any](v T) {
	fmt.Println(v)
	switch reflect.ValueOf(v).Kind() {
	case reflect.Ptr:
		fmt.Println("Its pointer")
	default:
		fmt.Println("Its not pointer")
	}
	fmt.Println()
}
