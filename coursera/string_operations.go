package main

import (
	"fmt"
	"strconv"
	"strings"
)

func toWeirdCase(str string) string {
	// Your code here and happy coding!
	out := make([]string, len(str))
	count := 0
	for i, val := range str {
		if val == ' ' || val >= '0' && val <= '9' {
			count = 0
			out[i] = string(val)
			continue
		}
		if count % 2 == 1 {
			out[i] = strings.ToLower(string(val))
		} else {
			out[i] = strings.ToUpper(string(val))
		}
		count++
	}
	return strings.Join(out, "")
}

func main() {
	{
		str := make([]string, 10)
		str[1] = "1"
		fmt.Println(str)
		result := strings.Join(str, "")
		fmt.Println(result)
	}
	{
		str := " "
		fmt.Printf("%v\n", str[0])
		s := strconv.Itoa(123)
		fmt.Printf("%v\n", s)
		s = toWeirdCase("AbC DeF")
		fmt.Printf("%v\n", s)
	}
	{
		c := '#'
		fmt.Printf("default value: %v\t type: %[1]T\t binary: %[1]b\t decimal: %[1]d\n", c)
	}
}