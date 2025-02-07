package func_args

import (
	"fmt"
)

type Data struct {
	num int
	str string
}

func CheckArgsInFn() {
	myData := &Data{}
	myStr := "default"

	defer func() {
		input(myData, myStr)
	}()

	myData.num = 100
	myData.str = "hello world"
	myStr = "new str"
}

func input(myData *Data, myStr string) {
	fmt.Println(myData, myStr)
}
