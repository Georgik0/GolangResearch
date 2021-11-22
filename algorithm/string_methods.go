package main

import (
	"fmt"
	"strings"
)

func concatString(a, b string) string {
	var new_string strings.Builder
	new_string.Grow(len(a) + len(b))
	new_string.WriteString(a)
	new_string.WriteString("45")
	return new_string.String()
}

func main() {
	var s1, s2 string
	s1 = "123"
	s2 = "456"
	str := concatString(s1, s2)
	fmt.Printf("str = %v\t len = %v\t str[4] = %v\n", str, len(str), str[4])
}
