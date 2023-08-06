package cases

import "fmt"

type A struct {
	s *string
}

func Test2() {
	a := &A{}

	if a.s == nil {
		fmt.Println("s is NIL")
	} else {
		fmt.Println("s isn't NIL")
	}
}
