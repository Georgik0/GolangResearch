package ptr_on_slice_element

import "fmt"

func Check() {
	arr := []int{0, 1, 2}

	num := &arr[2]
	arr[2] = 22
	fmt.Println("Before append after change to 22", *num)

	arr = append(arr, 3, 4, 5, 6, 7)
	arr[2] = 222
	fmt.Println("After append and change to 222", *num)
}
