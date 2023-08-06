package test3

import "fmt"

type data struct {
	M map[string]string
}

func Check() {
	d := data{}

	if d.M == nil {
		fmt.Println("It's nil")
	} else {
		d.M["test"] = "test"
	}

	fmt.Println(len(d.M), " map: ", d.M)
}
