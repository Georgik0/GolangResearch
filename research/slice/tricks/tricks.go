package tricks

import "fmt"

func ResearchCopy() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{0, 11, 22, 33, 44, 55, 66, 77, 88}
	c := []int{0}
	d := make([]int, 3, 3)

	//copy(a[8:], a[8+1:])
	//a = a[:len(a)-1]

	//copy(c[0:], c[0+1:])
	//c = c[:len(c)-1]

	//copy(d[3:], d[3+1:])
	//d = d[:len(d)-1]

	a = DeleteElementFromSlice(a, 4)
	b = DeleteElementFromSlice(b, 0)
	c = DeleteElementFromSlice(c, 0)
	d = DeleteElementFromSlice(d, 3)

	fmt.Println("a: ", a, " ", cap(a))
	fmt.Println("b: ", b, " ", cap(b))
	fmt.Println("c: ", c, " ", cap(c))
	fmt.Println("d: ", d, " ", cap(d))
}

func DeleteElementFromSlice[T any](arr []T, idx int) []T {
	if len(arr) == 0 {
		return []T{}
	}

	if int(idx) >= len(arr) {
		newArr := make([]T, len(arr))
		copy(newArr, arr)
		return newArr
	}

	newArr := make([]T, len(arr)-1)

	if int(idx) == len(arr)-1 {
		copy(newArr, arr[:len(arr)-1])
		return newArr
	}

	copy(newArr[:idx], arr[:idx])
	copy(newArr[idx:], arr[idx+1:])

	return newArr
}

func DelWithCopy[T any](arr []T, idx int) []T {
	//arr := []string{"1", "2", "3"}
	return arr[:idx+copy(arr[idx:], arr[idx+1:])]
}
