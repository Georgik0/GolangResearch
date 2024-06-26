package solution_2

func WordBreak(s string, wordDict []string) bool {
	wordDictMap := make(map[string]struct{}, len(wordDict))

	for _, w := range wordDict {
		wordDictMap[w] = struct{}{}
	}

	result := make([]bool, len(s)+1)
	result[0] = true

	for i := 0; i < len(s)-1; i++ {
		for j := len(s) - 1; j >= i+1; j-- {
			if _, ok := wordDictMap[s[i:j+1]]; ok {
				if result[i] {
					result[j+1] = true
				}
			}
		}
	}

	return result[len(s)]
}
