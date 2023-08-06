package append_test

func AppendInCycle() {
	tmp := map[int]int8{1: 1}

	count := 2
	for k, v := range tmp {
		tmp[count] = 1
		count++
		println(k, v)
	}
}
