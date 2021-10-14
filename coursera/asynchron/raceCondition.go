package main

import (
	"fmt"
	"sync"
)

func main() {
	dict := make(map[int]int)
	mu := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func(counters map[int]int, th int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mu.Lock()
				counters[th * 10 + j]++
				mu.Unlock()
			}
		}(dict, i, mu)
	}
	fmt.Scanln()
	mu.Lock()
	fmt.Println("dictionary result", dict)
	mu.Unlock()
}
