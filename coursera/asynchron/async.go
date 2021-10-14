package main

import "fmt"

func main() {
	in := make (chan int)

	/* chan<- int - в канал можно только писать; <-chan int - из канала можно только читать */
	go func(out chan<- int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("Before", i)
			out <- i
			fmt.Println("After", i)
		}
		close(out)
		fmt.Println("Finish")
	}(in)

	for i := range in {
		fmt.Println("\tget:", i)
	}
}
