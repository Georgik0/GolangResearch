package share_variable

import (
	"fmt"
	"sync"
	"time"
)

func Check() {
	a := "aaaaa"

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(a)
		wg.Done()
	}()

	a = "bbbbb"
	wg.Wait()
}

func Check2() {
	var num *int
	var result *int

	num = getNum(10)
	fmt.Println("num 1;", num)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		num = getNum(20)
		fmt.Println("result 2;", num)
		wg.Done()
	}()

	wg.Wait()
	result = num

	fmt.Println("result 3;", result)
	fmt.Println("num", num)
}

func getNum(b int) *int {
	a := b
	return &a
}
