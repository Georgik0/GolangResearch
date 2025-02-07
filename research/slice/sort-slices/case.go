package sortslices

import (
	"fmt"
	"sort"
)

func CaseSortNotCoppiedSlice() {
	sl := []string{"5", "4", "3", "2", "1"}
	newSL := sl

	sort.Strings(newSL)

	fmt.Println(sl)
}
