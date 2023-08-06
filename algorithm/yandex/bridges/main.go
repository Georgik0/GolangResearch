package main

import (
	"fmt"
	"os"
)

const noRestrictions = 1000000007

func main() {
	var numberOfBridges int
	fmt.Fscan(os.Stdin, &numberOfBridges)
	G_i_j := make([][]int64, numberOfBridges, numberOfBridges)
	Result_i_j := make([][]int64, numberOfBridges, numberOfBridges)

	for i := 0; i < numberOfBridges; i++ {
		G_i_j[i] = make([]int64, numberOfBridges, numberOfBridges)
		Result_i_j[i] = make([]int64, numberOfBridges, numberOfBridges)

		for j := 0; j < numberOfBridges; j++ {
			fmt.Fscan(os.Stdin, &G_i_j[i][j])
		}
	}

	for i := range G_i_j {

		for j := range G_i_j[i] {

			if i == j {
				Result_i_j[i][j] = noRestrictions
				continue
			}

		}

	}

	for _, v := range G_i_j {
		fmt.Println(v)
	}
}
