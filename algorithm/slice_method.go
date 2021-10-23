package main

import "fmt"

type Vector []interface{}

func (v *Vector) Pushback(elem interface{})  {
	*v = append(*v, elem)
}

func (v *Vector) Erase(i int) {
	if i >= len(*v) {
		return
	}
	*v = append((*v)[:i], (*v)[(i + 1):]...)
}

func reversString(list []string) Vector {
	fmt.Println("input list: ", list)
	out := Vector{[]string{}}
	out.Pushback("1")
	out.Pushback("2")
	out.Pushback("3")
	fmt.Printf("%v\t%[1]T\n", out)
	return out
}

func main() {
	reversString([]string{"test"})

	a := Vector{1, 2, 3}
	b := a
	b[1] = 9
	fmt.Printf("%v\t%[1]T\n", a)
	fmt.Printf("%v\t%[1]T\n", b)
	c := Vector{int16(11), int16(12), int16(13)}

}