package main

import (
	"fmt"
	"sync"

	map_concurrency "GolangResearch/research/concurrency/map-concurrency"
)

func main() {
	//case1.Check()
	//test1()
	map_concurrency.Check()
}

func test1() {
	wg := sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		fmt.Println("i до сапуска горутины: ", i)
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
