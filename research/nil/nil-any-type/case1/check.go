package case1

import (
	"fmt"
)

type Custom struct{}

func (c *Custom) F() {
	fmt.Println("c: ", c)
}

func myNil() *Custom {
	return nil
}

func Check() {
	c := myNil()
	if c != nil {
		fmt.Println("c is not NIL")
	}

	if c == nil {
		fmt.Println("c is NIL")
	}
}
