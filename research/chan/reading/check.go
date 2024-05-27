package reading

import (
	"fmt"
	"time"
)

func Check() {
	ch := make(chan struct{})

	go func() {
		ch <- struct{}{}
	}()

	time.Sleep(1 * time.Second)
	checkRead(ch)
	fmt.Println("finished")
}

func checkRead(ch chan struct{}) {
	select {
	case <-ch:
		fmt.Println("read chan successful")
	default:
		fmt.Println("skip reading")
	}
}
