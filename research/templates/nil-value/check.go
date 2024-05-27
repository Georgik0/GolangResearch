package nil_value

import (
	"fmt"
	"reflect"
)

type data struct {
	a int
}

func Check() {
	val := fReturnNilValue[*int]()
	fmt.Printf("%T %v\n", val, val)

	fmt.Printf("%T %v\n", example(), *example())

	f1 := fReturnDefaultZeroValue[*int]
	a := 5
	fmt.Printf("%T %v\n", f1(&a), f1(&a))
}

func fReturnDefaultZeroValue[T any](arg T) T {
	var zeroValue T

	argIsNil := reflect.ValueOf(arg).IsNil()
	zeroValueIsNil := reflect.ValueOf(zeroValue).IsNil()
	fmt.Printf("In value==nil --> %v value=%v\n", zeroValueIsNil, zeroValue)
	fmt.Printf("In arg==nil --> %v arg=%v\n", argIsNil, arg)

	return zeroValue
}

func fReturnNilValue[T any]() T {
	var zero T

	return zero
}

func example() *int {
	a := 0
	return &a
}
