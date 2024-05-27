package main

import "strings"

type obj struct {
	arg1, arg2, arg3, arg4 string
}

func (o *obj) Info() string {
	return strings.Join([]string{o.arg1, o.arg2, o.arg3, o.arg4}, "-")
}

// Что лучше использовать объект или анонимную функцию?
func main() {
	var arg1, arg2, arg3, arg4 string

	j := func() string {
		return strings.Join([]string{arg1, arg2, arg3, arg4}, "-")
	}

	j()
}
