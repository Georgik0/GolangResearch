package main

/*
j > i > 0
prefixSums[j] - prefixSums[i] = k
*/

func subarraySum(nums []int, k int) int {
	prefixSums := make(map[int]int, len(nums))
	allSums := 0
	currentPrefixSum := 0
	prefixSums[0] = 1

	for _, num := range nums {
		currentPrefixSum += num

		if count, exist := prefixSums[currentPrefixSum-k]; exist {
			allSums += count
		}

		prefixSums[currentPrefixSum]++
	}

	return allSums
}
