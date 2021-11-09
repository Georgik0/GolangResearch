package main

import (
	"fmt"
	"strings"
)

type set map[string]bool

/*func convert_to_slice(arr set) []int32 {
	convert_slice := []int32{}
	for key, _ := range arr {
		convert_slice = append(convert_slice, key)
	}
	return convert_slice
}*/

func recursive_string_permutation(s string) set {
	var out set = make(set)
	if (len(s) <= 1) {
		out[s] = true
		return out
	}
	lastChar := string(s[len(s) - 1])
	prevStr := s[:(len(s) - 1)]
	recursiveSet := recursive_string_permutation(prevStr)
	for strInSet := range recursiveSet {
		for pos := range prevStr {
			result := []string{strInSet[:pos], lastChar, strInSet[pos:]}
			p := strings.Join(result, "")

			out[p] = true

			//if pos == 1 || pos == 0 {
			//	fmt.Printf("i = %v: %v\t%v\t%v\n", pos, result, strInSet, lastChar)
			//}
			if pos == 0 {
				result := []string{strInSet, lastChar}
				p := strings.Join(result,"")
				out[p] = true
			}
		}
	}
	return out
}

func main() {
	out := recursive_string_permutation("1234")
	fmt.Println(out)
}
