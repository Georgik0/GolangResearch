package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	fd,err := os.Open("./data")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
}
