package main

import (
	"fmt"
	"sync"
)

func main() {
	researchConcurrentWriteInMap()
}

func researchConcurrentWriteInMap() {
	m := make(map[int]struct{})
	n := 10

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for j := range []int{1, 2, 3, 4, 5} {
				m[j] = struct{}{}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(m)
}
