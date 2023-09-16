package read_non_existent_element

import "fmt"

func Check() {
	m := make(map[int]int)

	m[1] = 1
	fmt.Println(m[2])
}
