package main

import "fmt"

func main() {
	/*gr := make(map[int][]int)
	fmt.Println("gr: ", gr)
	gr[0] = append(gr[0], 0)
	gr[1] = append(gr[1], 1)
	fmt.Println("After append\ngr: ", gr)
	change_map1(gr)
	fmt.Println("After change_map\ngr: ", gr)

	m2 := make(map[int]int)
	change_map2(m2)
	fmt.Println(m2)*/

	//equateSlice()
	emptyMap()
}

func emptyMap() {
	m := make(map[int]string)

	fmt.Println(m[1])
	_, ok := m[1]
	fmt.Println(ok)
}

func change_map1(gr map[int][]int) {
	gr[0] = append(gr[0], 0)
	gr[1] = append(gr[1], 1)
}

func change_map2(m map[int]int) {
	m[1] = 1
	m[2] = 22
	m[3] = 333
}

func equateSlice() {
	a := make([]int, 4, 4)
	for i := 0; i < 4; i++ {
		a[i] = i + 1
	}
	fmt.Println("a = ", a)

	b := make([]int, 2, 2)
	for i := 0; i < 2; i++ {
		b[i] = (i + 1) * 10
	}
	fmt.Println("b = ", b)

	c := make([]int, 8, 8)
	for i := 0; i < 8; i++ {
		c[i] = (i + 1) * 11
	}
	fmt.Println("c = ", c)
	a = b
	fmt.Println("a = b; a = ", a)
	a = c
	fmt.Println("a = c; a = ", a)
}
