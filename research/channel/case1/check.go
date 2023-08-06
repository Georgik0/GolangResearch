package case1

import "fmt"

func Check() {
	ch := make(chan int)

	// deadlock !!!
	for val := range ch {
		fmt.Println(val)
	}
}
