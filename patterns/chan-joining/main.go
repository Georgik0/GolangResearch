package main

import (
	"fmt"
	"sync"
)

func main() {
	chansSlice := make([]chan int, 5)

	for i := 0; i < len(chansSlice); i++ {
		chansSlice[i] = make(chan int)

		go func(ch chan<- int, idx int) {
			defer close(ch)

			fmt.Printf("start goroutine #%v\n", idx)

			for k := 3; k > 0; k-- {
				ch <- k * idx
			}
		}(chansSlice[i], i)
	}

	chanResult := ChanJoin(chansSlice...)
	for res := range chanResult {
		fmt.Println(res)
	}

	fmt.Println("finish")
}

func ChanJoin(in ...chan int) <-chan int {
	out := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(in))

	for _, pipe := range in {
		go func(ch <-chan int) {
			defer wg.Done()

			for v := range ch {
				out <- v
			}
		}(pipe)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
