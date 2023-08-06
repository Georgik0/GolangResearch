package main

import (
	"fmt"
	"regexp"
)

func main() {
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
		"典体",
	}

	for _, s := range data {
		//if checkWordRegexp(s) != checkWord(s) {
		//	fmt.Println("bad: ", "|"+s+"|")
		//}
		fmt.Println(checkWordRegexp(s), "   |"+s+"|")
	}
}

var exactlyOneWordPattern = regexp.MustCompile(`^\b\w+\b$`)

func checkWordRegexp(text string) bool {
	return exactlyOneWordPattern.MatchString(text)
}

func A_Z(c uint8) bool {
	return c >= 'A' && c <= 'Z'
}

func isNumber(c uint8) bool {
	return c >= '0' && c <= '9'
}

func a_z(c uint8) bool {
	return c >= 'a' && c <= 'z'
}

func correctSymbol(c uint8) bool {
	return A_Z(c) || a_z(c) || isNumber(c) || c == '_'
}

func checkWord(text string) bool {
	if text == "" {
		return false
	}

	if !correctSymbol(text[0]) {
		return false
	}

	for i := 0; i < len(text); i++ {
		if !correctSymbol(text[i]) {
			return false
		}
	}

	return true
}
