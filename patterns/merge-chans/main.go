package main

import "sync"

func main() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)

	go func() {
		for i := 0; i <= 3; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for i := 4; i <= 7; i++ {
			c2 <- i
		}
		close(c2)
	}()

	go func() {
		for i := 8; i <= 10; i++ {
			c3 <- i
		}
		close(c3)
	}()

	mergeC := mergeChans(c1, c2, c3)

	for n := range mergeC {
		println(n)
	}
}

func mergeChans(chs ...chan int) chan int {
	out := make(chan int)

	wg := sync.WaitGroup{}
	for _, ch := range chs {
		ch := ch

		wg.Add(1)

		go func() {
			for chanElement := range ch {
				out <- chanElement
			}

			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
