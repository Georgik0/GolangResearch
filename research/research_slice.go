package main

import "fmt"

func main() {
	//mystring(3)
	//fmt.Println(mystring(3))
	//fmt.Println(new_version(3))
	//s1, s2 := checkNilSlice()
	//fmt.Println(s1, s2)
	//fmt.Println(len(s1), len(s2))
	checkSeg()
}

func checkSeg() {
	arr := []int{1}

	a := arr[:1]
	fmt.Println(a)
}

func checkNilSlice() ([]int, []int) {
	zeroSl := []int(nil)
	return zeroSl, nil
}

func check_copy() {
	a := []string{"111", "222", "333", "444", "555"}

	fmt.Println(a)
	copy(a[1+1:], a[1:])
	fmt.Println(a)
}

var columNames = []string{"Query", "Checked", "AliasFor", "IsAdult", "CatalogName", "Glued", "ImagesUrl"}

func new_version(levels int) []string {
	result := make([]string, 0, 20) // why 20? 8 system fields + 8 max depth at 04.22 + some buff =)
	result = append(result, "ID")
	for i := 0; i < levels; i++ {
		result = append(result, fmt.Sprintf("%d уровень", i))
	}

	return append(result, columNames...)
}

func mystring(levels int) []string {
	result := []string{"ID", "Query", "Checked", "AliasFor", "IsAdult", "CatalogName", "Glued", "ImagesUrl"}
	for i := 1; i <= levels; i++ {
		result = append(result, "")
		copy(result[i+1:], result[i:])
		result[i] = fmt.Sprintf("%d уровень", i-1)
		//fmt.Println("i = ", i, result)
	}
	return result
}
