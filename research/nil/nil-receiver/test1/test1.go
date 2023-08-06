package test1

import "fmt"

type Rec struct {
	A int
}

func (r *Rec) Check() {
	fmt.Println(r.A)
}
