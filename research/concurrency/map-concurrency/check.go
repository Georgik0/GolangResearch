package map_concurrency

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func Check() {
	m := make(map[int]int)
	for i := 0; i <= 100; i++ {
		m[i] = i
	}

	runtime.GOMAXPROCS(4)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(iGo int) {
			mu.Lock()
			defer mu.Unlock()

			time.Sleep(time.Duration(rand.Intn(100000)))
			delete(m, iGo)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
