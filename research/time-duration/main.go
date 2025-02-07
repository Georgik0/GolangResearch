package main

import (
	"fmt"
	"time"
)

func main() {
	researchSleep()
	//checkParsing()
	//checkDuration()
}

func researchSleep() {
	start := time.Now()
	time.Sleep(time.Millisecond * 50)
	end := time.Since(start)

	fmt.Println(end.Milliseconds())
}

func checkParsing() {
	start := time.Now()

	fStart := start.Format(time.RFC3339)
	fStartNano := start.Format(time.RFC3339Nano)
	fmt.Printf("fStart: %s\nfStartNano: %s\n", fStart, fStartNano)

	parsedStart, err := time.Parse(time.RFC3339Nano, fStart)

	time.Sleep(time.Millisecond * 100)
	duration := time.Since(parsedStart)

	fmt.Printf("duration_ms: %v\nerr: %v", duration.Milliseconds(), err)
}

func checkDuration() {
	start := time.Now().Format(time.RFC3339)

	time.Sleep(time.Millisecond * 100)

	parsedStart, _ := time.Parse(time.RFC3339, start)
	duration := time.Since(parsedStart)

	fmt.Printf("duration_ms: %v\n", duration.Seconds())
}
