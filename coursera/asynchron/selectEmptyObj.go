package main

import (
	"fmt"
)

func main() {
	cancelCh := make(chan struct{})
	dataCh := make(chan int)

	go func(cancelCh chan struct{}, dataCh chan int) {
		val := 0
		for true {
			select {
			case <-cancelCh: // выполняется если мы что-то прочитали из канала cancelCh
				return
			case dataCh <- val: // выполняется в противном случае
				val++
			}
		}
	}(cancelCh, dataCh)

	for curVal := range dataCh {
		fmt.Println("read: ", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- struct{}{} // когда это произойдет, это значение прочитается в case <-cancelCh:
			break
		}
	}
}
