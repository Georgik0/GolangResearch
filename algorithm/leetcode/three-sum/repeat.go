package main

import (
	"slices"
)

func threeSumRepeat1(nums []int) [][]int {
	out := make([][]int, 0)

	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		println(i)
		left := i + 1
		right := len(nums) - 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				out = append(out, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			}
		}
	}

	return out
}
