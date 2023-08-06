package case1

import (
	"fmt"
	"time"
)

func Case() {
	fmt.Println("Start case ...")

	if true {
		defer func() { fmt.Println("defer is correct") }()
		time.Sleep(1)
	}

	fmt.Println("condition is finished")
}
