package main

import (
	"fmt"

	"GolangResearch/research/nil/nil_map/test1"
)

func main() {
	test1.Check()
	//test2.Test()
	//test1()
	//test3.Check()
	//test4.Check()
}

type MyMAP map[int]int

func localCheck() {
	checkMap(nil)
}

func checkMap(m MyMAP) {
	if v, ok := m[1]; true {
		fmt.Println(v, ok)
	}
}
