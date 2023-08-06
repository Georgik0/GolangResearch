package main

import "fmt"

func main() {
	sl := getNilSlice()

	fmt.Printf("sl: %v; len: %v; type: %T", sl, len(sl), sl)
}

func getNilSlice() []byte {
	return nil
}
