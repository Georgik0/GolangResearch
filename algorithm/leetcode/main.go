package main

import "fmt"

func main() {
	fmt.Println(countDigitOne(1))
}

func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}

	count := 0

	for i := 1; i <= n; i++ {
		count += getCountOfNumber(i)
	}

	return count
}

func getCountOfNumber(n int) int {
	count := 0

	for n != 0 {
		if n%10 == 1 {
			count++
		}
		n = n / 10
	}

	return count
}
