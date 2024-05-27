package research_slices

import "fmt"

func Research() {
	a := []int{1}
	b := a

	b = append(b, 2)

	info(b, "b")
	info(a[:1], "a")
}

func Research1() {
	b := make([]int, 3, 4)

	cb := append(b, 2)
	info(cb, "cb")
	info(b, "b")

	b[0] = 1
	info(cb, "cb")

	a := make([]int, 3, 3)
	ca := append(a, 2)
	info(ca, "ca")
	info(a, "a")

	a[0] = 1
	info(ca, "ca")
}

func Research2() {

}

func info(arr []int, arrName string) {
	fmt.Printf(arrName+"  arr: %v  cap=%d len=%d address=%p\n", arr, cap(arr), len(arr), arr)
}
