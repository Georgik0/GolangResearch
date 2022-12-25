package main

import (
	"context"
	"time"
)

func main() {

}

func researchContext() {
	ctx := context.Background()

	go func() {
		time.Sleep(1)
		select {
		case <-ctx.Done():
			print("ctx Done in: go func")
		}
	}()

}
