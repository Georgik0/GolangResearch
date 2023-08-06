package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	/* read line */
	{
		fd, err := os.Open("./data") // reader := bufio.NewReader(os.Stdin) - чтобы читать с Stdin
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
		scanner := bufio.NewScanner(fd)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Printf("%T\t%[1]v\n", text)
		}
	}
	/* read elem */
	{
		var a int
		fmt.Fscan(os.Stdin, &a)
		fmt.Printf("a: %v\n", a)
		fmt.Fscanf(os.Stdin, "%d", &a)
		fmt.Printf("a: %v\n", a)
	}
}
