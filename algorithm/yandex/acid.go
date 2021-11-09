package main

import (
	"fmt"
	"os"
)

func main() {
	var N int
	fmt.Fscan(os.Stdin, &N)
	if N <= 1 {
		fmt.Println(0)
		return
	}
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(os.Stdin, &arr[i])
	}
	first := arr[0]
	last := arr[len(arr) - 1]
	//fmt.Printf("first: %T\t%[1]v\n", first)
	//fmt.Printf("last: %T\t%[1]v\n", last)
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i - 1] {
			fmt.Println(-1)
			return
		}
	}
	//out := last - first
	fmt.Println(last - first)
}