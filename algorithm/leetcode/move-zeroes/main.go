package main

var nums = []int{0, 1, 0, 3, 12}

func main() {
	moveZeroes(nums)

	for _, v := range nums {
		println(v)
	}
}
