package main

import "fmt"

func getSlice(n int) *[]int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	return &a
}

func main() {
	a := getSlice(5)
	b := getSlice(3)
	c := getSlice(4)
	fmt.Println(a)

	sets := []([]int){}
	sets = append(sets, *a)
	sets = append(sets, *b)
	sets = append(sets, *c)
	fmt.Println("sets:", sets)
	c = a
	fmt.Println("c:", c)
	fmt.Println("sets:", sets)
}
