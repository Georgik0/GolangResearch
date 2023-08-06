package check

import "fmt"

var a = 1

func Check2() {
	afterInc := IncAfterDefer()
	fmt.Println("afterInc: ", afterInc)
}

func IncAfterDefer() int {
	defer Inc()

	return a
}

func Inc() {
	a += 1
}

func Show(a int) {
	fmt.Println("In return a: ", a)
}
