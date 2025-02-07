package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	go_retry "github.com/sethvargo/go-retry"
)

func main() {
	var counter = 0

	err := go_retry.Constant(context.Background(), time.Second, func(ctx context.Context) error {
		counter += 1

		if counter < 10 {
			fmt.Println("[<10] retrying, counter =", counter)

			return go_retry.RetryableError(errors.New("want retry"))
		}

		fmt.Println("[==10] error is nil, counter =", counter)

		return nil
	})

	if err != nil {
		fmt.Println("finish error:", err)
	}
}
