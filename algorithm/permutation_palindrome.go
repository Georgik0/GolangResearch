package main

import "fmt"

func PermutationPalindrome(in string) bool {
	check := map[int32]int{}
	for _, i := range in {
		_, ok := check[i]
		if ok {
			delete(check, i)
		} else {
			check[i]++
		}
	}
	if len(check) <= 1 {
		return true
	}
	return false
}

func main() {
	tests := []struct {
		in       string
		expected bool
	}{
		{"", true},
		{"c", true},
		{"cc", true},
		{"ca", false},
		{"civic", true},
		{"ivicc", true},
		{"civil", false},
		{"livci", false},
	}
	for _, val := range tests {
		if PermutationPalindrome(val.in) != val.expected {
			fmt.Printf("Error\n")
		}
	}
}
