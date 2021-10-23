package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

type meeting struct {
	start int
	end int
}

func MergeMeeting(meetings []meeting) []meeting {
	fmt.Println("Before sort: ", meetings)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].start < meetings[j].start
	})

	result := []meeting{}
	for idx, val := range meetings {
		if idx == len(meetings) - 1 {
			break
		}
		if meetings[idx].end >= meetings[idx + 1].start {

		}
	}
}
