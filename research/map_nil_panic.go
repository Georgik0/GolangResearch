package main

import "fmt"

type checkMap struct {
	kek map[string]string
	num int
}

func main() {
	chM := nilMap()
	fmt.Println(chM)
	if chM.kek["123"] == "123" {
		fmt.Println("OK")
	}
	fmt.Println(chM.kek["test"])
}

func nilMap() checkMap {
	chM := checkMap{}
	return chM
}
