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
	err error
}

func F() data {
	d := data{}

	start := time.Now()
	defer func() {
		d.end = time.Since(start)
		fmt.Printf("in F() d.end = %v; d.err = %v\n", d.end, d.err)
	}()

	time.Sleep(2 * time.Second)
	d.err = fmt.Errorf("new error")

	fmt.Println("out of F()")

	return d
}
