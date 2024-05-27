package main

import return_from_func "GolangResearch/research/nil/nil-interfaces/return-from-func"

func main() {
	//research()
	return_from_func.Research()
}

func research() {
	d := &data{
		a: 5,
	}

	var dd *data

	checkNil(d)
	checkNil(dd)
	checkNil(nil)
}

type Provider interface {
	Info()
	IsNil() bool
}

type data struct {
	a int
}

func (d *data) Info() {
	println(d.a)
}

func (d *data) IsNil() bool {
	return d == nil
}

func checkNil(pr Provider) {
	println(pr.IsNil())
	pr.Info()

	switch pr.(type) {
	case *data:
		println("true")
	}

	if pr == nil {
		println("nil")
		return
	}

	pr.Info()
}
