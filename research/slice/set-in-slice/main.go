package main

import "fmt"

func main() {
	sl := make([]int, 3, 4)

	copy(sl[3:], []int{2})

	fmt.Println(sl, "  ", sl[3])
}
