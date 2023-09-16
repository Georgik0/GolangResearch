package main

import "fmt"

func getSubsets(a []int, k int) {
	n := len(a)
	init := make([]int, n)
	for i := 0; i < n; i++ {
		if i < k {
			init[i] = 1
		} else {
			init[i] = 0
		}
	}

	//allSubset := []([]int){}

	for true {
		current := 0
		unit := 0

		//currentSet := make([]int, k)

		for current < n && (init[current] == 1 || unit == 0) {
			if init[current] == 1 {
				unit++
			}
			current++
		}

		if current == n {
			break
		}
		init[current] = 1

		for i := 0; i < current; i++ {
			if i < unit-1 {
				init[i] = 1
			} else {
				init[i] = 0
			}
		}
		fmt.Println(init)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	getSubsets(a, 3)
}

/*
	Необходимо вывести все подмножества слайса a, размером k
	Пример
	a := []int{1,2,3,4,5,6}
	getSubsets(a, 3)

	Вывод
	[1 1 0 1 0 0]
	[1 0 1 1 0 0]
	[0 1 1 1 0 0]
	[1 1 0 0 1 0]
	[1 0 1 0 1 0]
	[0 1 1 0 1 0]
	[1 0 0 1 1 0]
	[0 1 0 1 1 0]
	[0 0 1 1 1 0]
	[1 1 0 0 0 1]
	[1 0 1 0 0 1]
	[0 1 1 0 0 1]
	[1 0 0 1 0 1]
	[0 1 0 1 0 1]
	[0 0 1 1 0 1]
	[1 0 0 0 1 1]
	[0 1 0 0 1 1]
	[0 0 1 0 1 1]
	[0 0 0 1 1 1]
*/
