package test4

import "fmt"

type getter interface {
	get()
}

type mChecker map[string]getter

func Check() {
	m := make(mChecker)

	v, ok := m["test"]
	fmt.Println(v, ok)
	//v.get()
}
