package detect_type

import (
	"fmt"
	"reflect"
)

type MyData struct{}

func (d *MyData) Info() {}

type MyDataI interface {
	Info()
}

func Research() {
	var ptr *MyData

	dI := MyDataI(nil)
	fmt.Println(reflect.TypeOf(dI))
	Check(ptr)
}

func Check(d MyDataI) {
	fmt.Println(reflect.TypeOf(d).String())
}
