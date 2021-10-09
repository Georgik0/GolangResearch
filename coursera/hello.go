package main

import (
	"fmt"
	"math/big"
	"strings"
)

type Person struct {
	name string
	age int
}

type MySlice []int

func (sl *MySlice) add(i int) {
	*sl = append(*sl, i)
}

func (p *Person) SetName(name string) {
	p.name = name
}

func main() {

	/* Объявление переменных */
	fmt.Println("\n--------------------Объявление переменных------------------------")
	{
		var str_1 string = "Option 1"
		var str_2 = "Option 2"
		str_3 := "Option_3"

		fmt.Printf("1: %v   2: %v   3: %v\n", str_1, str_2, str_3)
	}

	/* Пробуем range */
	fmt.Println("\n--------------------range------------------------")
	{
		s1 := "x~@#$%"
		for _, c := range s1 {
			fmt.Printf("c: %5v  %3[1]c  %3[1]T\n", c)
		}
		fmt.Println("")
	}

	/* Конвертация типов */
	fmt.Println("\n--------------------Конвертация типов------------------------")
	{
		var f = 123.123
		var i64 = int64(f)
		var i32 = int32(i64)
		var i16 = int16(i32)
		var i8 = int8(i16)
		var i = int(i8)
		fmt.Printf("float = %v  i64 = %v  i32 = %v  i16 = %v  i8 = %v  i = %v", f, i64, i32, i16, i8, i)
	}


	/* Пробуем смайлики */
	var smile rune = 1285515
	fmt.Println(smile)
	fmt.Printf("smile: %c\n", smile)
	//fmt.Printf("smile: \1285515\n")

	fmt.Println("Hello")

	var str = "New string"
	var out = strings.Compare(str, "w s")
	fmt.Printf("Результат compare: %v\n", out)

	out = strings.Compare(str, "New string")
	fmt.Printf("Результат compare: %v\n\n", out)

	var data interface{}
	data = 'a'
	switch data_type := data.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Printf("type = %T\n", data_type)
		fmt.Println("value = ", data_type)
	}

	var dist = new(big.Int)
	fmt.Printf("dist = %v   type = %[1]T\n", dist)

	/* Массивы */
	fmt.Println("\n--------------------Массивы------------------------")
	{
		const size = 4
		var arr1 [5]int
		fmt.Println(arr1, "длина", len(arr1))

		arr2 := [...]int{1, 2, 3, 4}
		fmt.Println(arr2, "длина", len(arr2))

		var arr3 [size]int
		fmt.Println(arr3, "длина", len(arr3))
	}

	/* Слайсы */
	fmt.Println("\n--------------------Слайсы------------------------")
	{
		var arr []int
		fmt.Println("Пустой слайс", arr, "size = ", len(arr))
		arr = append(arr, 10)
		fmt.Println("Добавили элемент", arr, "size = ", len(arr))
		fmt.Println("Длина внутреннего массива cap = ", cap(arr))
		arr = append(arr, 20)
		fmt.Println("Добавили второй элемент", arr, "size = ", len(arr))
		fmt.Println("Длина внутреннего массива cap = ", cap(arr))
		arr = append(arr, 30)
		fmt.Println("Добавили третий элемент", arr, "size = ", len(arr))
		fmt.Println("Длина внутреннего массива cap = ", cap(arr))
		arr = append(arr, 40)
		fmt.Println("Добавили четвертый элемент", arr, "size = ", len(arr))
		fmt.Println("Длина внутреннего массива cap = ", cap(arr))
		arr = append(arr, 50)
		arr = append(arr, 60)
		fmt.Println("Добавили пятый шестой элемент", arr, "size = ", len(arr))
		fmt.Println("Длина внутреннего массива cap = ", cap(arr))

		fmt.Println("Добавим слайс к слайсу")
		add := []int{
			1, 2, 3,
		}
		 arr = append(arr, add...)
		 fmt.Println("Добавили add", arr, "cap =", cap(arr))

		 var arr2 [][]int
		 arr2 = append(arr2, arr)
		 fmt.Println("Двумерный слайс", arr2)

		 slice3 := make([]int, 10)
		 slice3 = append(slice3, 1)
		 fmt.Println("Слайс нужной длины", slice3)

		 slice_copy := slice3
		 slice_copy[1] = 123
		 fmt.Println("Изменили значение по ссылке на слайс\nslice = ", slice3, "\nslice_copy = ", slice_copy)

		 var slice_copy3 []int
		 copy(slice_copy3, slice3)
		 fmt.Println("slice_copy3 = ", slice_copy3, "\nslice3 = ", slice3)
		 correct_copy3 := make([]int, len(slice3), len(slice3))
		 copy(correct_copy3, slice3)
		 fmt.Println("correct_copy3 = ", correct_copy3)
	}

	/* map */
	fmt.Println("\n--------------------Map------------------------")
	{
		var mm map[string]string
		fmt.Println("uninitialized map", mm)

		mm2 := map[string]string{}
		mm2["test"] = "ok"
		fmt.Println(mm2)

		var mm3 = make(map[string]string)
		mm3["one"] = "1"
		fmt.Println("mm3 = ", mm3)
		elem := mm3["something"]
		fmt.Println("non-existent elem: ", elem)

		var value, check_value = mm3["something"]
		fmt.Println("value = ", value, "\ncheck_value = ", check_value)

		delete(mm3, "test")
		_, exist_check := mm3["test"]
		fmt.Println("mm3[\"test\"] is ", exist_check)
	}

	/* if */
	fmt.Println("\n--------------------if------------------------")
	{
		mm := map[string]string{
			"one": "1",
			"two": "2",
			"three": "3",
		}
		if value, check := mm["two"]; check {
			fmt.Println("value: ", value, check)
		} else {
			fmt.Println("non-existent element", value, check)
		}
	}

	/* for */
	fmt.Println("\n--------------------for range------------------------")
	{
		sl := []int{10, 20, 30, 40}
		fmt.Println("range slice")
		for idx, val := range sl {
			fmt.Println(idx, ":", val)
		}

		mm := map[string]int{
			"one":1,
			"two":2,
		}
		fmt.Println("range map")
		for key, value := range mm{
			fmt.Println(key, ":", value)
		}
	}

	/* switsch case*/
	fmt.Println("\n--------------------switch case------------------------")
	{
		mm := make(map[string]float64)
		mm["one"] = 1.0
		mm["two"] = 2
		switch {
		case mm["one"] == 1.0:
			fmt.Println("value = ", mm["one"])
		}
	}

	/* анонимные функции */
	fmt.Println("\n--------------------anonim func------------------------")
	{
		type str_func func(in string)
		print_str := func(str string) {
			fmt.Println(str)
		}
		view_in := func(in_func str_func, slice []string) {
			in_func(slice[0])
		}
		arr := []string{"test func"}
		view_in(print_str, arr)
	}

	/* struct и методы */
	fmt.Println("\n--------------------struct и методы------------------------")
	{
		Bob := Person {
			"Noname",
			30,
		}
		fmt.Println("Bob: ", Bob)
		Bob.SetName("Bob")
		fmt.Println("After SetName(\"Bob\"): ", Bob)
	}

	go_vector := MySlice{}
	go_vector.add(10)
	fmt.Println("vector: ", go_vector)
	var go_vector1 = make(MySlice, 10, 11)
	fmt.Printf("go_vector1: %v\ttyp: %[1]T\tcap: %v\tlen: %v\n", go_vector1, cap(go_vector1), len(go_vector1))
	go_vector1.add(123)
	fmt.Printf("go_vector1: %v\ttyp: %[1]T\tcap: %v\tlen: %v\n", go_vector1, cap(go_vector1), len(go_vector1))

}
