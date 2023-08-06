package check

import (
	"fmt"
	"time"
)

func Check() {
	d := F()
	fmt.Println("d := F()  -->  d.end = ", d.end)
}

type data struct {
	end time.Duration
}

func F() data {
	d := data{}

	start := time.Now()
	defer func() {
		d.end = time.Now().Sub(start)
		fmt.Println("in F() d.end = ", d.end)
	}()

	time.Sleep(2)
	fmt.Println("out of F()")

	return d
}
