package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fd,err := os.Open("./data")
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
