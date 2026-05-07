package main

import (
	"fmt"
)

var input = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

func main() {
	volume := trap(input)

	fmt.Println(volume)
}
