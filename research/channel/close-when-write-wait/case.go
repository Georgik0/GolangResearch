package close_when_write_wait

import (
	"fmt"
	"sync"
	"time"
)

func Research() {
	wg := sync.WaitGroup{}

	wg.Add(1)

	ch := make(chan struct{})
	exit := make(chan struct{})

	go func() {
		fmt.Println("Start write")

		select {
		case ch <- struct{}{}:
			fmt.Println("Write finished")
		case <-exit:
			fmt.Println("Exit")
		}

		wg.Done()
	}()

	time.Sleep(time.Second)
	go func() {
		fmt.Println("Close")

		close(ch)
	}()

	wg.Wait()
}
