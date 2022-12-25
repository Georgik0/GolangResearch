package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetCategoryID(path string) {
	rCatSeoID := regexp.MustCompile(`/category/(\w*-)*(\d+)(/|$)`)
	res := rCatSeoID.FindStringSubmatch(rCatSeoID.FindString(path))
	fmt.Println(res)
}

func main() {
	GetCategoryID("/modal/allFilters/category/elektronika-15500/?debug=1&text=%D1%81%D0%BC%D0%B0%D1%80%D1%82%D1%84%D0%BE%D0%BD%D1%8B")

	id, _ := strconv.ParseInt("", 10, 64)
	fmt.Println(id)

	s := ""
	//s = s[:len(s)-1] //panic
	fmt.Println(s)
}
