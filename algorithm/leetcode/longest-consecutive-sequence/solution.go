package main

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		set[num] = struct{}{}
	}

	maxLength := 1
	for num := range set {
		localMax := 1

		if _, ok := set[num-1]; ok {
			// Значит num - не начало последовательности
			continue
		}

		for {
			num++
			if _, has := set[num]; has {
				localMax++
			} else {
				break
			}
		}

		if localMax > maxLength {
			maxLength = localMax
		}
	}

	return maxLength
}
