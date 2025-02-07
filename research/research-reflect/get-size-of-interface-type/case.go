package get_size_of_interface_type

import (
	"fmt"
	"reflect"
)

func Research() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arrb := []byte("test")
	var i interface{}

	_ = arr
	i = arrb

	fmt.Println(int(reflect.TypeOf(i).Size()))
}
