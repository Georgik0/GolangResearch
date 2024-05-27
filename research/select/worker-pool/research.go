package worker_pool

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type worker struct {
	read   chan struct{}
	cancel chan struct{}
	number int
	ctx    context.Context
}

func (w *worker) Start() {
	for {
		fmt.Println(strconv.Itoa(w.number) + " start")
		select {
		case <-w.read:
			fmt.Println(strconv.Itoa(w.number) + " read")
		case <-w.cancel:
			fmt.Println(strconv.Itoa(w.number) + " cancel")
			return
		case <-w.ctx.Done():
			fmt.Println(strconv.Itoa(w.number) + " ctx.Done()")
			return
		}
	}
}

func WorkerPoolStart() {
	readC := make(chan struct{})
	cancelC := make(chan struct{})
	ctx, cancelFunc := context.WithCancel(context.Background())

	for i := 0; i <= 9; i++ {
		newWorker := worker{
			read:   readC,
			cancel: cancelC,
			number: i,
			ctx:    ctx,
		}

		go newWorker.Start()
	}

	for i := 0; i <= 30; i++ {
		readC <- struct{}{}
	}

	//cancelC <- struct{}{}
	//cancelC <- struct{}{}

	cancelFunc()

	time.Sleep(time.Millisecond * 10000)
}
