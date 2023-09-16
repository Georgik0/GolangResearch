package main

import (
	"fmt"
	"strings"
)

// Написать функцию, которая принимает строку и возвращает все уникальные перестановки символов

type set map[string]struct{}

/*func convert_to_slice(arr set) []int32 {
	convert_slice := []int32{}
	for key, _ := range arr {
		convert_slice = append(convert_slice, key)
	}
	return convert_slice
}*/

func recursive_string_permutation(s string) set {
	var out = make(set)

	if len(s) <= 1 {
		out[s] = struct{}{}

		return out
	}

	lastChar := string(s[len(s)-1])
	prevStr := s[:(len(s) - 1)]
	recursiveSet := recursive_string_permutation(prevStr)

	for strInSet := range recursiveSet {
		for pos := range prevStr {
			result := []string{strInSet[:pos], lastChar, strInSet[pos:]}
			p := strings.Join(result, "")

			out[p] = struct{}{}

			if pos == 0 {
				result := []string{strInSet, lastChar}
				p := strings.Join(result, "")
				out[p] = struct{}{}
			}
		}
	}

	return out
}

func main() {
	out := recursive_string_permutation("1234")

	for k := range out {
		fmt.Println(k)
	}
}
