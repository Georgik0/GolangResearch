package case3

import (
	"sync"
)

var wg sync.WaitGroup

func Check() {
	ch := make(chan int)

	wg.Add(1)
	go startReader(ch)
	ch <- 1
	wg.Wait()
}

func startReader(ch chan int) {
	for i := range ch {
		println(i)
		ch <- i
	}
	wg.Done()
}
