package main

import (
	solution_1 "GolangResearch/algorithm/leetcode/word-break/solution-1"
	solution_2 "GolangResearch/algorithm/leetcode/word-break/solution-2"
)

func main() {
	println(solution_2.WordBreak("leetcode", []string{"leet", "code"}))
	println(solution_2.WordBreak("a", []string{"a"}))
	//s := "1234"
	//println(s[0:3])
	println(solution_1.WordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa"}))
}
