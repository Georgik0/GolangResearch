package main

import (
	"testing"
)

func TestInflight(t *testing.T) {
	test := []struct{
		in1 []int
		in2 int
		expected bool
	} {
		{[]int{}, 1, false},
		{[]int{0}, 1, false},
		{[]int{1}, 1, false},
		{[]int{0, 1}, 1, true},
		{[]int{1, 1}, 2, true},
		{[]int{2, 3, 4}, 6, true},
		{[]int{2, 3, 4}, 8, false},
	}
	for _, val := range test {
		if out := Inflight(val.in1, val.in2); out != val.expected {
			t.Errorf("Error:\nout: %v\nexpected: %v\n", out, val)
		}
	}
}
