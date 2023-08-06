package test2

import "fmt"

type MapOfSlices map[string][]*int

func Test() {
	m := make(MapOfSlices)
	n := 10

	m["1"] = append(m["1"], &n)
	fmt.Println(m)
}
