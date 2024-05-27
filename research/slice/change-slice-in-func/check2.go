package change_slice_in_func

import "fmt"

const (
	arrCopyName = "arrCopy"
	arrName     = "arr"
)

func Check2() {
	//arr := []int{0, 1, 2}
	arr := make([]int, 0, 4)
	arr = append(arr, 0)
	arr = append(arr, 1)
	arr = append(arr, 2)

	arrCopy := arr

	fmt.Printf("arr: %p, %v, %v\n", arr, cap(arr), len(arr))
	fmt.Printf("arrCopy: %p, %v, %v\n", arrCopy, cap(arrCopy), len(arrCopy))

	changeSlice(arr, arrName)
	changeSlice(arrCopy, arrCopyName)

	fmt.Printf("arrCopy: %p, %v, %v, arr - %v\n", arrCopy, cap(arrCopy), len(arrCopy), arrCopy)
}

func changeSlice(argArr []int, name string) {
	fmt.Printf(
		"in func (%s): %p, %v, %v; arr - %v\n",
		name,
		argArr,
		cap(argArr),
		len(argArr),
		argArr,
	)

	if name == arrCopyName {
		argArr[2] = 0
		argArr = append(argArr, 4)

		fmt.Printf(
			"in func (%s) -- argArr[2] = 0 -- append -- info: %p, %v, %v; arr - %v\n",
			name,
			argArr,
			cap(argArr),
			len(argArr),
			argArr,
		)
	}
}

/* Output
arr: 0x14000126000, 4, 3
arrCopy: 0x14000126000, 4, 3
in func (arr): 0x14000126000, 4, 3; arr - [0 1 2]
in func (arrCopy): 0x14000126000, 4, 3; arr - [0 1 2]
in func (arrCopy) -- argArr[2] = 0 -- append -- info: 0x14000126000, 4, 4; arr - [0 1 0 4]
arrCopy: 0x14000126000, 4, 3, arr - [0 1 0]
*/
