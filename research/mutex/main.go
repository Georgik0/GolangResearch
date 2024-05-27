package main

import "sync"

type data struct {
	mu sync.RWMutex
}

func main() {
	d := data{}

	d.mu.RLock()
	d.mu.RUnlock()
}
