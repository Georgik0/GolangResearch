package main

import "fmt"

func Inflight(list []int, flightLength int) bool {
	//fmt.Println("list", list)
	if len(list) == 0 {
		if flightLength == 0 {
			return true
		} else {
			return false
		}
	}
	for i, val1 := range list[:(len(list)-1)] {
		for _, val2 := range list[(i+1):] {
			fmt.Println(val1, val2)
			if	val2 + val1 == flightLength {
				return true
			}
			}
		}
	return false
}

/*func fillFlight(movieLengths []int, flightLength int) bool {
	movies := map[int]int{}

	for _, v := range movieLengths {
		// return true if we've seen the movie. else add the movie in the
		// watched list with a count 1.
		matchLength := flightLength - v
		if _, ok := movies[matchLength]; ok {
			return true
		}

		movies[v] = 1
	}

	return false
}*/

func main() {
	test := []struct{
		in1 []int
		in2 int
		expected bool
	} {
		{[]int{}, 1, false},
		{[]int{0}, 1, false},
		{[]int{1}, 1, false},
		{[]int{0, 1}, 1, true},
		{[]int{1, 1}, 2, true},
		{[]int{2, 3, 4}, 6, true},
		{[]int{2, 3, 4}, 8, false},
	}
	for _, val := range test {
		if out := Inflight(val.in1, val.in2); out != val.expected {
			fmt.Printf("Error:\nout: %v\nexpected: %v\n", out, val)
		}
	}
}
