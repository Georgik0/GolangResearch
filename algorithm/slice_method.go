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

func swapElemSlice(i, j int, vector *Vector) {
	if i >= len(*vector) || j >= len(*vector) {
		panic("Segmentation fault")
	}
	tmp := (*vector)[i]
	(*vector)[i] = (*vector)[j]
	(*vector)[j] = tmp
}

func reversString(list []string) []string {
	middle := len(list) / 2
	out := make([]string, len(list))
	for i := 0; i < middle; i++{
		out[i] = list[len(list) - i - 1]
		out[len(list) - i - 1] = list[i]
	}
	if len(list) % 2 != 0 {
		out[middle] = list[middle]
	}
	return out
}

func main() {
	result_vec := reversString([]string{"1","2","3","4"})
	fmt.Printf("result: %v  %[1]T\n", result_vec)
	result_vec = reversString([]string{})
	fmt.Printf("result: %v  %[1]T\n", result_vec)
	result_vec = reversString([]string{"1"})
	fmt.Printf("result: %v  %[1]T\n", result_vec)
	result_vec = reversString([]string{"1", "2", "3"})
	fmt.Printf("result: %v  %[1]T\n", result_vec)

	a := Vector{1, 2, 3}
	b := a
	b[1] = 9
	fmt.Printf("%v\t%[1]T\n", a)
	fmt.Printf("%v\t%[1]T\n", b)
	c := Vector{int16(11), int16(12), int16(13)}
	fmt.Printf("%v\t%[1]T\n", c)
	c.Pushback(123)
	fmt.Printf("%v\t%[1]T\ttype c[0]: %T\n", c, c[3])
}