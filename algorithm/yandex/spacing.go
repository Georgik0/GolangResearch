package main

import (
	"fmt"
	"math"
)

var min float64 = math.Inf(1)

func dist(a_i int, S []int) {
	var summ float64 = 0
	for _, elem := range S {
		summ += math.Abs(float64(a_i - elem))
	}
	if summ	< min {
		min = summ
	}
}

func genericS(arr []int, index int, k int) {
	init := []int{}
	for i, _ := range init {
		if i < k {
			init[i] = 1
		} else {
			init[i] = 0
		}
	}
	n := len(arr)
	for {
		cur_id := 0
		unit := 0
		for i := 0; i < k; i++ {

		}
	}
}

func main() {
	var N, k int
	fmt.Scanf("%d %d", &N, &k)
	arr := []int{}
	var elem int
	for i := 0; i < N; i++ {
		fmt.Scan(&elem)
		arr = append(arr, elem)
	}
	for i := 0; i < N; i++ {

	}
}
