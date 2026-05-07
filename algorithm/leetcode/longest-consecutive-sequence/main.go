package main

import (
	"fmt"
)

var in = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

func main() {
	res := longestConsecutive(in)

	fmt.Println(res)
}
