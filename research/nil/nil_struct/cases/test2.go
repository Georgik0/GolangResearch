package cases

import "fmt"

type A struct {
	s *string
}

func Test2() {
	a := new(A)

	if a.s == nil {
		fmt.Println("s is NIL")
	} else {
		fmt.Println("s isn't NIL")
	}

	checkNilA(nil)
}

func checkNilA(obj *A) {
	if obj.s == nil {
		fmt.Println("s is NIL")
	} else {
		fmt.Println("s isn't NIL")
	}
}
