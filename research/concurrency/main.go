package main

import (
	"fmt"
	"sync"

	research_errgroup "GolangResearch/research/concurrency/research-errgroup"
)

func main() {
	//case1.Check()
	//test1()
	//map_concurrency.Check()
	//slice_concurrency.Check()
	//share_variable.Check2()
	research_errgroup.Case()
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
