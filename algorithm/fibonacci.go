package main

import (
	"fmt"
	"sync"
)

func fibonacci(n int) int {
	var vector []int = []int{0, 1}
	for i := 2; i <= n; i++ {
		vector = append(vector, vector[i - 1] + vector[i - 2])
	}
	return vector[n]
}

func fibonacci_recursiv(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	fib := fibonacci_recursiv(n-1) + fibonacci_recursiv(n-2)
	return fib
}

func recurs(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	N := fibonacci_recursiv(n)
	fmt.Println("recurs N = ", N)
}

func dinamic(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	N := fibonacci(n)
	fmt.Println("dinamic N = ", N)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go recurs(40, wg)
	wg.Add(1)
	go dinamic(40, wg)
	wg.Wait()
}
