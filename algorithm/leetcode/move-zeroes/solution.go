package main

func moveZeroes(nums []int) {
	left := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		nums[left] = nums[i]
		left++
	}

	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}
}
