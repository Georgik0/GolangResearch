package main

import "fmt"

func main() {
	a := []string{"0", "1", "2"}
	b := a
	b[1] = "11"
	a[0] = "00"

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}

func help() {
	locs := []int64{}
	var AreaID int64

	{
		var locations []int64

		if AreaID != 0 {
			locations = make([]int64, len(locs)+1)
			locations[0] = AreaID
			copy(locations[1:], locs)
		}
	}

	{
		locations := make([]int64, 0, len(locs)+1)

		if AreaID != 0 {
			locations = append(locations, AreaID)
		}

		locations = append(locations, locs...)
	}
}
