package main

import "fmt"

type researchStruct struct {
	a *int
}

type researchStruct2 struct {
	a int
	b string
	c *float64
}

func main() {
	//research_struct()
	//research_struct_ptr()
	research_slice_ptr()
}

func research_slice_ptr() {
	a := make([]int, 0)
	fmt.Println("after change:	len =", len(a), "cap =", cap(a), "	a =", a)
	change_slice(&a)
	fmt.Println("after change:	len =", len(a), "cap =", cap(a), "	a =", a)
}

func change_slice(a *[]int) {
	for i := 0; i < 10; i++ {
		*a = append(*a, i)
	}
	fmt.Println("after change:	len =", len(*a), "cap =", cap(*a), "	a =", a)
}

func research_struct_ptr() {
	s := struct {
		a int
		b string
		c *float64
	}{a: 1000, b: "hahaha"}
	fmt.Println(s, s.b)

	var d *researchStruct2 = nil
	if d != nil {
		fmt.Println(d.a)
	}
}

func research_struct() {
	var i int = 10
	cs := &researchStruct{&i}
	fmt.Println("Before change: ", *cs.a)
	changeStruct(*cs)
	fmt.Println(*cs.a)
	changeStructPtr(cs)
	fmt.Println(*cs.a)
}

func changeStruct(cs researchStruct) {
	var i int = 20
	cs.a = &i
}

func changeStructPtr(cs *researchStruct) {
	var i int = 30
	cs.a = &i
}
