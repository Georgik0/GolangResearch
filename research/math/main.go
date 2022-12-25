package main

import "fmt"

func main() {
	checkType()
}

func checkType() {
	a := int64(100) + 1
	fmt.Printf("type - %T\n", a)
}

func checkMaxInt() {
	var a int64
	a = 9223372036854775807
	fmt.Println(a)
	a += 1
	fmt.Println(a)
}

func mod() {
	a := 11
	fmt.Println(a % 10)
}
