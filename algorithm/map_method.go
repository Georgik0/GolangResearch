package main

import "fmt"

func main() {
	{
		var dr map[string]int = make(map[string]int)
		dr["1"] = 1
		fmt.Println(dr)
	}
	{
		dr := map[string]int{
			"one":1,
			"two":2,
		}
		fmt.Println(dr, &dr)
		fmt.Printf("%v\t%[1]T\n", &dr)
	}

}
