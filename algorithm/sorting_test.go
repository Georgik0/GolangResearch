package main

import (
	"fmt"
	"sort"
	"testing"
)

type meeting struct {
	start int
	end int
}

func EqualSlice(a []meeting, b []meeting) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	for i, val_a := range a {
		if i >= len(b) {
			return false
		}
		if val_a.start != b[i].start || val_a.end != b[i].end {
			return false
		}
	}
	return true
}

func TestMergeMeetings(t *testing.T) {
	tests := []struct {
		in []meeting
		expected []meeting
	}{
		{[]meeting{}, []meeting{}},
		{[]meeting{{1, 2}}, []meeting{{1, 2}}},
		{[]meeting{{1, 2}, {2, 3}}, []meeting{{1, 3}}},
		{[]meeting{{1, 5}, {2, 3}}, []meeting{{1, 5}}},
		{[]meeting{{1, 2}, {4, 5}}, []meeting{{1, 2}, {4, 5}}},
		{[]meeting{{1, 5}, {2, 3}, {4, 5}}, []meeting{{1, 5}}},
		{[]meeting{{1, 2}, {2, 3}, {4, 5}}, []meeting{{1, 3}, {4, 5}}},
		{[]meeting{{1, 6}, {2, 3}, {4, 5}}, []meeting{{1, 6}}},
		{[]meeting{{4, 5}, {2, 3}, {1, 6}}, []meeting{{1, 6}}},
	}
	for _, tt := range tests {
		result := mergeMeeting(tt.in)
		if EqualSlice(result, tt.expected) == false {
			t.Errorf("fail\nresult = %v\nexpected = %v\n", result, tt.expected)
		}
	}
}

func mergeMeeting(meetings []meeting) []meeting {
	fmt.Println("Before sort: ", meetings)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].start < meetings[j].start
	})

	result := []meeting{}
	for idx := 0; idx < len(meetings); idx++ {
		if idx == len(meetings) - 1 {
			break
		}
		if meetings[idx].end >= meetings[idx + 1].start {
			if meetings[idx].end > meetings[idx + 1].end {
				result = append(result, meetings[idx])
			} else {
				elem := meeting{start: meetings[idx].start, end: meetings[idx + 1].end}
				result = append(result, elem)
			}
			idx++;
		} else {
			result = append(result, meetings[idx])
		}
	}
	fmt.Println("After sort: ", result)
	return result
}
