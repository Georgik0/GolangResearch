package change_slice_in_func

import "fmt"

type data struct {
	name string
	age  int
}

var arrGlobal = []data{
	{
		name: "1",
		age:  1,
	},
	{
		name: "2",
		age:  2,
	},
	{
		name: "3",
		age:  3,
	},
}

func Check() {
	fmt.Println("before changePublicSlice public arrGlobal:", arrGlobal)
	changePublicSlice(arrGlobal)
	fmt.Println("after changePublicSlice public arrGlobal:", arrGlobal)
}

func changePublicSlice(arrArg []data) {
	for _, v := range arrArg {
		v.name = "newName"
	}
	fmt.Println("in func changePublicSlice after 'for' slice:", arrArg)
}
