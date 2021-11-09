package main



func MakeChange(num []int, summ int) int {
	result := 0
	for _, val := range num {
		if summ - val < 0 {
			continue
		}
		if summ - val == 0 {
			result++
		}
		if summ - val > 0 {
			result += MakeChange(num, summ - val)
		}
	}
	return result
}
