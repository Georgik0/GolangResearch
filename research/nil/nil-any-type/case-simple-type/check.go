package case_simple_type

import "fmt"

func Check() {
	var aPtr *int

	//a := 10

	aPtr = nil
	fmt.Println(*aPtr)
}
