package main

import (
	"fmt"
	"time"
)

func main() {
	var zi interface{}

	zi = i1
	fmt.Println(zi)
}

func i1(a int, s string) float64 {
	return 0
}

func myTimer(f func(interface{}, interface{}) float64) {

}

func research2() {

}

func research1() {
	start := time.Now()

	for i := 1; i < 100000; {
		i++
	}

	time.Sleep(200000)
	endSince := time.Since(start)
	endSub := time.Now().Sub(start)

	fmt.Println("Since")
	fmt.Println("endSince int64", endSince)
	fmt.Println("endSub explicit int64", int64(endSince))
	checkDuration(endSince)
	fmt.Println("String: ", endSince.String())
	fmt.Println("Seconds: ", endSince.Seconds())
	fmt.Println("Milliseconds: ", endSince.Milliseconds())
	fmt.Println("Nanoseconds: ", endSince.Nanoseconds())

	fmt.Println("\nSub")
	fmt.Println("endSub int64", endSub)
	fmt.Println("endSub explicit int64", int64(endSub))
	checkDuration(endSub)
	fmt.Println("String: ", endSub.String())
	fmt.Println("Seconds: ", endSub.Seconds())
	fmt.Println("Milliseconds: ", endSub.Milliseconds())
	fmt.Println("Nanoseconds: ", endSub.Nanoseconds())

	fmt.Println("\nZero Duration")
	var td time.Duration
	checkDuration(td)
}

func checkDuration(t time.Duration) {
	fmt.Println("t: ", t)

	if t == 0 {
		fmt.Println("t == 0")
	}
}
