package main

func main() {

}

// Не решено
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{}, len(wordDict))

	for _, v := range wordDict {
		wordMap[v] = struct{}{}
	}

	for _, char := range s {
		buf := make([]byte, 0)
		buf = append(buf, uint8(char)) // вместо того чтобы хранить массивы байтов, нужно оперировать индексами

		if _, ok := wordMap[string(buf)]; ok {

		}
	}
}
