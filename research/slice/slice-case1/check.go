package slicecase1

import "fmt"

type Data struct {
	s string
}

func Check() {
	d := []Data{{"1"}, {"2"}, {"3"}}

	for _, v := range d {
		v.s = "new string"
	}

	fmt.Println(d)
}

func CheckNil() {
	var sl []int = nil

	if len(sl) == 0 {
		fmt.Println("NIL")
	}
}

func CheckLen0() {
	sl := make([]int, 0, 0)

	i := 1
	fmt.Println(sl[i-2])
}
