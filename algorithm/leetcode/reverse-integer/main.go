package main

import (
	"math"
)

func main() {
	//fmt.Printf("%v\n%v", math.MaxInt32, 1534236469)
	reverseV2(90000)
}

func reverseV1(x int) int {
	arrOfReversX := make([]int, 0, 2)

	sign := false
	if x < 0 {
		x = -x
		sign = true
	}

	for x > 0 {
		arrOfReversX = append(arrOfReversX, x%10)
		x /= 10
	}

	decade := 1
	revers := 0
	for i := len(arrOfReversX) - 1; i >= 0; i-- {
		revers += decade * arrOfReversX[i]
		decade *= 10
	}

	if sign {
		revers = -revers
	}

	if revers > math.MaxInt32 || revers < math.MinInt32 {
		return 0
	}

	return revers
}

func reverseV2(x int) int {
	sign := false
	if x < 0 {
		x = -x
		sign = true
	}

	tmpX := x
	maxDecade := 1
	for tmpX >= 10 {
		tmpX /= 10
		maxDecade *= 10
	}

	revers := 0
	for x > 0 {
		num := x % 10
		revers += maxDecade * num
		x /= 10
		maxDecade /= 10
	}

	if sign {
		revers = -revers
	}

	if revers > math.MaxInt32 || revers < math.MinInt32 {
		return 0
	}

	return revers
}
