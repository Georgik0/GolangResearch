package main

import "fmt"

func main() {
	//fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstringV3("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	startSubStringIndex := 0
	i := 0
	max := 0
	tmp := map[uint8]struct{}{}

	for i < len(s) {
		if _, ok := tmp[s[i]]; ok {
			if startSubStringIndex < i {
				delete(tmp, s[startSubStringIndex])
				startSubStringIndex++
				continue
			}
		}

		tmp[s[i]] = struct{}{}

		if i-startSubStringIndex > max-1 {
			max = i - startSubStringIndex + 1
		}

		i++
	}

	return max
}

func lengthOfLongestSubstringV2(s string) int {
	if len(s) == 1 {
		return 1
	}

	startMax := 0
	endMax := 0
	tmpStart := 0
	tmpIndex := 0

	for i := 1; i < len(s); i++ {
		tmpIndex = -1

		for j := tmpStart; j < i; j++ {
			if s[j] == s[i] {
				tmpIndex = j
				break
			}
		}

		if tmpIndex >= 0 {
			if (i - tmpStart - tmpIndex) > (endMax - startMax) {
				startMax = tmpIndex
				endMax = i
			}

			tmpStart = tmpIndex + 1
		} else {
			if (i - tmpStart + 1) > (endMax - startMax) {
				startMax = tmpStart
				endMax = i + 1
			}
		}
	}

	return endMax - startMax
}

func lengthOfLongestSubstringV3(s string) int {
	if len(s) == 1 {
		return 1
	}

	left, right, maxLen := 0, 0, 0

	tmp := make(map[uint8]int)

	for right < len(s) {
		rC := s[right]

		if rID, ok := tmp[rC]; ok && rID >= left {
			left = rID + 1
		}

		tmp[rC] = right

		if maxLen < right-left {
			maxLen = right - left
		}

		right++
	}

	return maxLen + 1
}
