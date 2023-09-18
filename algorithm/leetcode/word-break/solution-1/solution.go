package solution_1

// WordBreak Не проходит по времени
func WordBreak(s string, wordDict []string) bool {
	tmp := make(map[int32]struct{})

	for _, word := range wordDict {
		for _, c := range word {
			tmp[c] = struct{}{}
		}
	}

	for _, c := range s {
		if _, ok := tmp[c]; !ok {
			return false
		}
	}

	return execCompare(s, wordDict)
}

func checkPrefixWord(s, checkingWord string) bool {
	lenWord := len(checkingWord)

	if len(s) < lenWord {
		return false
	}

	return s[:lenWord] == checkingWord
}

func execCompare(s string, wordDict []string) bool {
	for _, word := range wordDict {

		if checkPrefixWord(s, word) {
			if s == word {
				return true
			}

			if execCompare(s[len(word):], wordDict) {
				return true
			}
		}

	}

	return false
}
