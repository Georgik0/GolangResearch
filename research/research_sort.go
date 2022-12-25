package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	arr := make([]int64, 0)

	_ = json.Unmarshal([]byte(`[1,2,3,4,5]`), &arr)
	fmt.Println(arr)
}
