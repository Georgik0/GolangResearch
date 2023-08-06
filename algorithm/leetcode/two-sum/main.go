package main

type data struct {
	val   int
	index int
}

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int, len(nums))

	for i, v := range nums {
		if val, ok := tmp[target-v]; ok {
			return []int{i, val}
		}

		tmp[v] = i
	}

	lenNums := len(nums)
	honest := lenNums%2 == 0
	if honest {
		for i, j := 0, lenNums-1; i < lenNums/2; i, j = i+1, j-1 {
			if val, ok := tmp[target-nums[i]]; ok {
				return []int{i, val}
			}
			tmp[nums[i]] = i

			if val, ok := tmp[target-nums[j]]; ok {
				return []int{j, val}
			}
			tmp[nums[j]] = j
		}
	} else {
		border := (lenNums - 1) / 2
		for i, j := 0, lenNums-1; i < border; i, j = i+1, j-1 {
			if val, ok := tmp[target-nums[i]]; ok {
				return []int{i, val}
			}
			tmp[nums[i]] = i

			if val, ok := tmp[target-nums[j]]; ok {
				return []int{j, val}
			}
			tmp[nums[j]] = j
		}

		if val, ok := tmp[target-nums[border]]; ok {
			return []int{border, val}
		}
	}

	return nil
}

func main() {
	twoSum([]int{3, 3}, 6)
}
