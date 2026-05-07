package main

func trap(height []int) int {
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	maxLocalLeft := 0
	for i, h := range height {
		maxLeft[i] = maxLocalLeft
		if h >= maxLocalLeft {
			maxLocalLeft = h
		}
	}

	maxLocalRight := 0
	for i := len(height) - 1; i >= 0; i-- {
		maxRight[i] = maxLocalRight
		if height[i] >= maxLocalRight {
			maxLocalRight = height[i]
		}
	}

	//fmt.Println(maxLeft, maxRight)

	sumVolume := 0
	for i, h := range height {
		vol := min(maxLeft[i], maxRight[i]) - h
		if vol > 0 {
			sumVolume += vol
		}
	}

	return sumVolume
}
