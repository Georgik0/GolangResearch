package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	fmt.Println("Start WaitGroup Done")
	wg.Done()
	fmt.Println("Finish WaitGroup Done")

	fmt.Println("Start WaitGroup Wait")
	wg.Wait()
	fmt.Println("Finish WaitGroup Wait")
}
