package main

import (
	"fmt"
	"strings"
)

func WordCloud(in string) map[string]int {
	out := map[string]int{}
	separate := strings.Split(in, " ")
	for _, val := range separate {
		out[val]++
	}
	return out
}

func main() {
	getMap := WordCloud("Cliff finished his cake and paid the bill. Bill finished his cake at the edge of the cliff.")
	for _, i := range getMap {
		fmt.Printf("%T\t%[1]v\n", i)
	}
}
