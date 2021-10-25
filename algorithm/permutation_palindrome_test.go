package main

import (
	"testing"
)

func PermutationPalindromeTest(t *testing.T) {
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
			t.Errorf("Error:\n")
		}
	}
}