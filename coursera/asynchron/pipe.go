package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pipe := make(chan int, 2)

	go func(out <-chan int) {
		for i := range out {
			fmt.Println("Read: ", i)
		}
		fmt.Println("Finish")
	}(pipe)

	for i := 0; i < 10; i++ {
		if i < 3 {
			fmt.Println("Before write: ", i)
			pipe <- i
			fmt.Println("After write: ", i)
		}
		if i == 6 {
			close(pipe)
		}
	}
	user := bufio.NewReader(os.Stdin)
	line, _, _ := user.ReadLine()
	fmt.Println(line)
}
