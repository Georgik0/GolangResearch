package case2

import "fmt"

type Data struct {
	s   string
	arr []int
	n   *int
}

func (d Data) Constructor(newD Data) {
	d = newD
}

func (d Data) SetS(newS string) {
	d.s = newS
	d.Show("SetS")
}

func (d Data) SetN(newN int) {
	*d.n = newN
	d.Show("SetN")
}

func (d Data) SetArr(nums ...int) {
	d.arr[0] = 1
	d.arr = append(d.arr, nums...)
	d.Show("SetArr: ")
}

func (d Data) Show(comment string) {
	if comment == "" {
		comment = "default show: "
	} else {
		comment += ": "
	}
	fmt.Println(comment, d.s, d.arr, *d.n)
}

func Case() {
	n := 1
	d := Data{s: "some string", arr: make([]int, 1, 10), n: &n}
	testD := Data{s: "NEW DATA"}

	d.SetS("1")
	d.Show("after SetS")
	fmt.Println()

	d.Constructor(testD)
	d.Show("after Constructor")
	fmt.Println()

	d.SetArr(2, 3, 4, 5, 6, 7, 8)
	d.Show("after SetArr")
	fmt.Println()

	d.SetN(100)
	d.Show("after SetN")
}
