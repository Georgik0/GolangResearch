package check

import (
	"fmt"
	"time"
)

func TimeDuration() {
	start := time.Now()
	defer PrintTime(
		time.Since(start),
	)
	defer func() {
		PrintTime(
			time.Since(start),
		)
	}()

	time.Sleep(3 * time.Second)
}

func PrintTime(duration time.Duration) {
	fmt.Println(duration.Seconds())
}
