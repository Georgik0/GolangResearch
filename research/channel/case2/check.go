package case2

func Check() {
	ch := make(chan int, 1)

	for val := range ch {
		println(val)
	}
}
