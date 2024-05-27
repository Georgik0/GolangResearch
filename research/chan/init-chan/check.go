package init_chan

import (
	"fmt"
	"runtime"
	"time"
)

// Нельзя допускать записи и чтения из nil канала

func Check() {
	runtime.GOMAXPROCS(2)
	var c chan int
	//c := make(chan int)

	go func() {
		fmt.Println(<-c)
		fmt.Println("thread 1 is ended")
	}()

	c <- 5

	fmt.Println("thread main")
	time.Sleep(2 * time.Second)
}
