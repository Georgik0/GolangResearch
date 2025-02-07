package trafficlimiter

import (
	"strconv"
	"sync"
	"testing"
)

func TestTrafficLimitersDispatcher(t *testing.T) {
	t.Parallel()

	dispatcher := NewDispatcher()
	defer dispatcher.CancelAll()

	const (
		writeThreadsNumber = 50
		readThreadsNumber  = 50
		rwThreadsNumber    = 100
	)

	wg := sync.WaitGroup{}

	// проверка на конкурентную запись
	wg.Add(writeThreadsNumber)
	for i := 0; i < writeThreadsNumber; i++ {
		go func(key int) {
			limiter := StartNewLimiter(0, 0)

			dispatcher.Add(limiterKey(strconv.Itoa(key)), limiter)

			wg.Done()
		}(i)
	}

	wg.Wait()

	// проверка на конкурентное чтение
	wg.Add(readThreadsNumber)
	for i := 0; i < readThreadsNumber; i++ {
		go func(key int) {
			_ = dispatcher.GetLimiterByKey(limiterKey(strconv.Itoa(key)))

			wg.Done()
		}(i)
	}

	wg.Wait()

	// проверка на конкурентные чтение/запись
	wg.Add(rwThreadsNumber)
	for i := 0; i < rwThreadsNumber; i++ {
		go func(key int) {
			limiter := StartNewLimiter(0, 0)

			dispatcher.Add(limiterKey(strconv.Itoa(key)), limiter)

			_ = dispatcher.GetLimiterByKey(limiterKey(strconv.Itoa(key)))

			wg.Done()
		}(i)
	}

	wg.Wait()

	dispatcher.CancelAll()
}
