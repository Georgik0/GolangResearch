package main

import (
	"testing"
)

type inTest struct {
	a []int
	b []int
}

func Equal(a, b []int) bool {
	if	len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMergeSortedArrays(t *testing.T) {
	in1 := inTest{
		a: []int{1, 3, 5},
		b: []int{2, 4, 6},
	}
	out1 := []int{1, 2, 3, 4, 5, 6}
	in2 := inTest{
		a: []int{},
		b: []int{4, 44, 444},
	}
	out2 := []int{4, 44, 444}
	in3 := inTest{
		a: []int{4, 44, 444},
		b: []int{},
	}
	out3 := []int{4, 44, 444}

	test := []struct{
		in inTest
		out []int
	}{
		{in1, out1},
		{in2, out2},
		{in3, out3},
	}
	for _, iter := range test {
		result := MergeSortedArrays(iter.in.a, iter.in.b)
		if Equal(result, iter.out) == false {
			t.Errorf("Not equal:\nresult: %v\nexpected:%v\n", result, iter.out)
		}
		t.Logf("Equal:\nresult: %v\nexpected:%v\n", result, iter.out)
	}
}

func MergeSortedArrays(a []int, b []int) []int {
	out := vectorInt{}
	len_a := len(a)
	len_b := len(b)

	ib := 0
	ia := 0
	for ia < len_a && ib < len_b {
		if a[ia] < b[ib] {
			out.Pushback(a[ia])
			ia++
		} else {
			out.Pushback(b[ib])
			ib++
		}
	}
	for _, v := range a[ia:] {
		out.Pushback(v)
	}
	for _, v := range b[ib:] {
		out.Pushback(v)
	}
	return getSliceInt(out)
}
