package case_with_slice_and_map

import (
	"fmt"
	"reflect"
)

type mySlice []int

func (s *mySlice) Info() {
	fmt.Println(*s)
}

func Case() {
	var sl mySlice

	fmt.Println(reflect.TypeOf(sl).Kind() == reflect.Pointer)
	//slNil := mySlice(nil)
	//
	//sl.Info()
	//slNil.Info()
}
