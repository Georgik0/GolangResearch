package main

import (
	"fmt"
	"os"
)

/*
1234

1
12				13		14
123		124		134
1234

2
23 24
234

3
34

*/

func getSets(a []int, current []int) []([]int) {
	sets := []([]int){}
	for i, val := range a {
		send := append(current, val)
		sets = append(sets, send)
		fmt.Println("sets add val", sets)
		sets = append(sets, getSets(a[i+1:], send)...)
		fmt.Println("sets add other sets", sets, "\n")
	}
	return sets
}

func main() {
	var N int
	fmt.Fscan(os.Stdin, &N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(os.Stdin, &a[i])
	}
	fmt.Println(a)
	sets := getSets(a, []int{})
	fmt.Println(sets)
	//test(a)
	//fmt.Println(a)
}
