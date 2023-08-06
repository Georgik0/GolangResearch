package main

import "fmt"

type myMap map[int]int

func (mm *myMap) Message() string {
	fmt.Println("Hello")

	return "return Hello"
}

type messager interface {
	Message() string
}

func main() {
	//var m myMap = nil

	//checkNil(&m)
	checkNil(nil)
}

func checkNil(m messager) {
	fmt.Println(fmt.Sprintf("%v", m.Message()))
}

func checkNilOfValue(m myMap) {
	fmt.Println(fmt.Sprintf("%v", m.Message()))
}
