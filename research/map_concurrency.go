package main

import (
	"fmt"
	"sync"
)

func test_1(wg *sync.WaitGroup, mu sync.Mutex) {
	dict := make(map[int]int)
	fmt.Println("Before", dict)

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			dict[i] = i
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("After", dict, "\n")
}

func test_2(wg *sync.WaitGroup, mu sync.Mutex) {
	dict := make(map[int]int)
	fmt.Println("Before", dict)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			dict[i] = i
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("After", dict, "\n")
}

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	test_1(&wg, mu)
	test_2(&wg, mu)
}
