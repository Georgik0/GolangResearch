package main

import "unsafe"

type provider interface {
	Get()
	Set(int) error
}

func main() {
	unsafe.Sizeof(provider)
}
