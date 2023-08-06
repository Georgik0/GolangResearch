package main

import "sort"

func main() {
	//threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0})
	//threeSum([]int{0, 0, 0})
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}

func threeSum(nums []int) [][]int {
	results := make([][]int, 0)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		currentValue := nums[i]

		for left < right {
			sum := -(nums[left] + nums[right])

			if currentValue > sum {
				right--
			} else if currentValue < sum {
				left++
			} else {
				result := []int{
					currentValue,
					nums[left],
					nums[right],
				}

				results = append(results, result)

				for left < right {
					if nums[left] != result[1] {
						break
					}
					left++
				}

				for left < right {
					if nums[right] != result[2] {
						break
					}
					right--
				}
			}

		}

		for i < len(nums)-1 {
			if nums[i] != nums[i+1] {
				break
			}
			i++
		}

	}

	return results
}
