package pointer_from_fn

import (
	"fmt"
)

type Data struct {
	num int
	str string
}

func Check() {
	var myData *Data

	defer func() {
		if myData == nil {
			fmt.Println("myData is nil")

			return
		}

		fmt.Println("myData: ", *myData)
	}()

	myData = getData()
	changeData(myData)
}

func getData() *Data {
	return &Data{
		num: 101,
		str: "create data",
	}
}

func changeData(d *Data) {
	d.num = 1
	d.str = "hello"
}
