package main

import "fmt"

/*
1234

1
12				13		14
123		124		134
1234

2
23 24
234

3
34

*/

/*func getSets(a []int, current []int) []([]int) {
	sets := []([]int){}
	for i, val := range a {
		send := append(current, val)
		sets = append(sets, send)
		fmt.Println("sets add val", sets)
		sets = append(sets, getSets(a[i+1:], send)...)
		fmt.Println("sets add other sets", sets, "\n")
	}
	return sets
}

func main() {
	var N int
	fmt.Fscan(os.Stdin, &N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(os.Stdin, &a[i])
	}
	fmt.Println(a)
	sets := getSets(a, []int{})
	fmt.Println(sets)
}*/

var all_sets = []([]int){}

func gen(a, current []int) {
	tmp := make([]int, len(current))
	copy(tmp, current)
	all_sets = append(all_sets, tmp)
	for id, val := range a {
		current = append(current, val)
		gen(a[id+1:], current)
		current = current[:len(current)-1]
	}
}

var all_perm = []([]int){}
var length int = -1

func permutation(a []int, chosen_id int) {
	if length == -1 {
		length = len(a)
	}
	for id, val := range a {
		tmp = append(tmp, val)
		if len(tmp) == length {
			all_perm = append(all_perm, tmp)
		}
	}
}

func main() {
	a := []int{1, 2, 3}

	fmt.Println("subsets:")
	gen(a, []int{})
	for _, val := range all_sets {
		fmt.Println(val)
	}

	fmt.Println("permutation:")
	permutation(a, []int{})
	for _, val := range all_perm {
		fmt.Println(val)
	}
}
