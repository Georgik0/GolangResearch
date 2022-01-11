package main

import "fmt"

func sl(a, b []int) []int {
	a_map := make(map[int]int)
	out := []int{}
	for _, val := range a {
		a_map[val]++
	}
	for _, val := range b {
		if _, ok := a_map[val]; ok {
			out = append(out, val)
			a_map[val]--
			if a_map[val] == 0 {
				delete(a_map, val)
			}
		}
	}
	return out
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 6, 6, 1}
	b := []int{0, 1, -2, -3, 4, 5, 5, 5, 5, 5, 6}
	result := sl(a, b)
	fmt.Println("Result:")
	for _, val := range result {
		fmt.Println("val = ", val)
	}
}
