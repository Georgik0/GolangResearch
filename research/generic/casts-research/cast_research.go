package casts_research

import "fmt"

func Check() {

}

func castGeneric[T any](arg any) {
	val, ok := arg.(*T)
	if !ok {
		fmt.Println(val)
	}
}
