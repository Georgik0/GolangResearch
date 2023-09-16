package main

import (
	"fmt"
)

func main() {
	index := 3
	wordDict := []string{"0", "1", "2", "3", "4", "5", "6"}

	newDict := make([]string, len(wordDict)-1)
	copy(newDict[:index], wordDict[:index])
	copy(newDict[index:], wordDict[index+1:])

	for i, v := range newDict {
		fmt.Printf("%v |%v|\n", i, v)
	}
	fmt.Println(newDict)

	//print(wordBreak("leetcode", []string{"leet", "code"}))
	print(wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
}

// Не решено
func wordBreak(s string, wordDict []string) bool {
	tmp := make(map[int32]struct{})

	for _, word := range wordDict {
		for _, c := range word {
			tmp[c] = struct{}{}
		}
	}

	for _, c := range s {
		if _, ok := tmp[c]; !ok {
			return false
		}
	}

	return execCompare(s, wordDict)
}

func checkPrefixWord(s, checkingWord string) bool {
	lenWord := len(checkingWord)

	if len(s) < lenWord {
		return false
	}

	return s[:lenWord] == checkingWord
}

func execCompare(s string, wordDict []string) bool {
	for _, word := range wordDict {

		if checkPrefixWord(s, word) {
			if s == word {
				return true
			}

			//newDict := make([]string, len(wordDict)-1)
			//copy(newDict[:index], wordDict[:index])
			//copy(newDict[index:], wordDict[index+1:])

			//newS := make([]byte, len(s)-len(word))
			//copy(newS, s[len(word):])

			if execCompare(s[len(word):], wordDict) {
				return true
			}
		}

	}

	return false
}
