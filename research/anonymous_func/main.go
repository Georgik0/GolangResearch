package main

import "fmt"

func main() {
	research1()
}

func modifier(f func()) {
	f()
}

func research1() {
	i := 1
	str := "str"

	f := func() {
		i += 3
		str += str
	}

	modifier(f)
	fmt.Println(i, str) // Out: 4 strstr
}
