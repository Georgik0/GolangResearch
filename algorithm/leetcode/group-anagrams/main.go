package main

import (
	"fmt"
)

var in = []string{"eat", "tea", "tan", "ate", "nat", "bat"}

func main() {
	res := groupAnagrams(in)

	for _, v := range res {
		fmt.Println(v)
	}
}
