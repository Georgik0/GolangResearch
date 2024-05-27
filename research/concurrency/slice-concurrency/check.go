package slice_concurrency

import "sync"

func Check() {
	sl := []int64{}

	wg := sync.WaitGroup{}
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			sl = append(sl, 1)
			wg.Done()
		}()
	}

	wg.Wait()

}
