package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {
	var str string
	for str != "123" {
		fmt.Scanf("%s\n", &str)
		fmt.Println("Out: ", str)
	}
}
