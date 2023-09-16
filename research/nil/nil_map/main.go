package main

import (
	"fmt"

	read_non_existent_element "GolangResearch/research/nil/nil_map/read-non-existent-element"
)

func main() {
	//test1.Check()
	//test2.Test()
	//test1()
	//test3.Check()
	//test4.Check()
	read_non_existent_element.Check()
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
