package cases

import "fmt"

func nilInterface() {
	m := make(map[int]interface{})

	a := 222
	s := "sadasdas"

	m[1] = a
	m[2] = nil
	m[3] = s

	switch m[1].(type) {
	case int:
		fmt.Println(m[1].(int))
	}

	if res, ok := m[1].(string); ok {
		fmt.Println(res)
	} else {
		fmt.Println("Not string")
	}

	var i interface{} = nil

	if res, ok := i.(int); ok {
		fmt.Println(res)
	} else {
		fmt.Println("Not int", res, ok)
	}
}

type testNilI interface {
	checker()
}

type testNil struct {
	s string
}

func (t *testNil) checker() {
	if t == nil {
		fmt.Println("t == nil - true")
		return
	}
	fmt.Println(t.s)
}

func initT() *testNil {
	return nil
}

func problemNIL(tI testNilI) {
	if tI == nil {
		fmt.Println("BEDA")
	}
	tI.checker()
}

func CheckTestNIL() {
	initT().checker()
}
