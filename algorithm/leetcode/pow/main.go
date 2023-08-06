package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(2, 10))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	count := 1
	tmp := make(map[int]float64)
	maxPower := 2
	offset := 0

	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	}

	result := x
	for true {
		if count*2 > n {
			offset = n - count
			break
		} else {
			count *= 2
			result *= result
			tmp[count] = result
			maxPower = count
		}
	}

	for offset > 1 {
		halfPower := maxPower / 2
		for halfPower > offset {
			halfPower /= 2
		}

		result *= tmp[halfPower]
		offset -= halfPower
		maxPower = halfPower
	}

	if offset == 1 {
		result *= x
	}

	if sign < 0 {
		result = 1 / result
	}

	return result
}

// 2 2 2 2 2	2 2 2	->	2^3 * 2^2

// 2^2  2^3  2^4  2^5
