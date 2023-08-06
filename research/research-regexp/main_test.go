package main

// go test -bench=.

import (
	"fmt"
	"testing"
)

const text = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111 111111 111111"

func BenchmarkCheckWordRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkWordRegexp(text)
	}
}

func BenchmarkCheckWordMy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkWord(text)
	}
}

func TestMyFunc(t *testing.T) {
	data := []string{
		"",
		"123",
		" 123",
		"123 123",
		"1",
		"1234567890qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM",
		"\\\\",
		" ",
		`фывапр`,
		`фывапр典体`,
	}

	for _, s := range data {
		//if checkWordRegexp(s) != checkWord(s) {
		//	fmt.Println("bad: ", "|"+s+"|")
		//}
		fmt.Println(checkWordRegexp(s), "   |"+s+"|")
	}

}
