package main

import "fmt"

func main() {
	//sl := getNilSlice()
	//
	//fmt.Printf("sl: %v; len: %v; type: %T", sl, len(sl), sl)

	checkNilArg(nil)
	checkNilArg([]string{})
}

func getNilSlice() []byte {
	return nil
}

func checkNilArg(arr []string) {
	fmt.Printf("arr: %v; len: %v; type: %T; cap: %v; arr==nil - %v\n", arr, len(arr), arr, cap(arr), arr == nil)
}
