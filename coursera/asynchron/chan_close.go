package main

import (
	"fmt"
)

func main() {
	pipe := make(chan int, 1)

	close(pipe)

	//pipe<- 1
	a := <-pipe
	fmt.Println(a)

	arr := make([]int, 2, 5)
	fmt.Println(arr, cap(arr))
	arr = append(arr, 1)
	fmt.Println(arr, cap(arr))

}
