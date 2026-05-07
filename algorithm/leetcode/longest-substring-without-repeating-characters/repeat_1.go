package main

func lengthOfLongestSubstring_1(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left, right, maxLength := 0, 0, 0

	unicChars := make(map[uint8]int)

	for right < len(s) {
		if _, ok := unicChars[s[right]]; ok {
			if left < unicChars[s[right]]+1 {
				left = unicChars[s[right]] + 1
			}
		}

		unicChars[s[right]] = right

		if right-left > maxLength {
			maxLength = right - left
		}

		right++
	}

	return maxLength + 1
}
