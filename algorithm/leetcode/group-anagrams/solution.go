package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		copyStrRune := []rune(str)
		slices.Sort(copyStrRune)

		sortedStr := string(copyStrRune)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	result := make([][]string, len(anagrams))
	i := 0
	for _, anagram := range anagrams {
		result[i] = anagram
		i++
	}

	return result
}
