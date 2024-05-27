package init_slice

import "fmt"

func Check() {
	sl := []int{
		1: 1,
		6: 4,
	}

	fmt.Printf("slice: %v; len: %v; cap: %v", sl, len(sl), cap(sl))
}
