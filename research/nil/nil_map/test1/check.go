package test1

import "fmt"

func Check() {
	var m map[int]int

	m = nil

	for k, v := range m {
		fmt.Printf("key: %v; value: %v ", k, v)
	}
}
