package main

import (
	"fmt"
	"context" // основная задача - отмена асинхронных операций
	"math/rand"
	"time"
)

func worker(ctx context.Context, workerNum int, out chan int) {
	waitTime := time.Duration(rand.Intn(100) + 10) * time.Millisecond
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done(): // обрабатывает сигнал finish()
		return
	case <-time.After(waitTime): // говорит о том что подошло время воркера
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}

func main() {
	ctx, finish := context.WithCancel(context.Background()) // ctx - контекст, который имеет функцию отмены finish()
	result := make(chan int, 1)

	for i := 0; i <= 10; i++ {
		go worker(ctx, i, result)
	}

	foundBu := <-result
	fmt.Println("result found by", foundBu)
	finish() // дает сигнал о том, что дальше продолжать не надо

	time.Sleep(time.Second)
}
