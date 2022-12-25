package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

var jsonMarshaler = jsoniter.ConfigCompatibleWithStandardLibrary

type someType struct {
	a int
	b bool
	s string
	l string
	c int64
}

type data struct {
	TestStringIntOtherName string `json:"testStringIntOtherName"`
	TestStringBool         bool   `json:"testStringBool"`
	Name1                  int    `json:"Name1"`
	Name2                  int    `json:"Name2"`
	Name3                  int    `json:"Name3"`
	TTTT                   bool   `json:"TTTT"`
	TestStringInt          int    `json:"TestStringInt"`
}

type dataSome struct {
	Name2                  int    `json:"Name2"`
	TestStringIntOtherName string `json:"testStringIntOtherName"`
	Name1                  int    `json:"Name1"`
	TestStringInt          int    `json:"TestStringInt"`
	Name3                  int    `json:"Name3"`
	TTTT                   bool   `json:"TTTT"`
	TestStringBool         bool   `json:"testStringBool"`
}

type dataMap map[string]string

type data1 struct {
	Num int
}

func main() {
	case1()
	//case2()
	//case3()
}

func case1() {
	src := map[string]interface{}{"TTTT": true, "TestStringBool": true, "Name2": 22, "name1": 11, "Name3": 33}
	srcData := data{
		TTTT:           true,
		TestStringBool: true,
		Name2:          22,
		Name1:          11,
		Name3:          33,
	}

	ByteMarshaler, _ := json.Marshal(src)
	fmt.Println(string(ByteMarshaler))
	ByteMarshalerData, _ := json.Marshal(srcData)
	fmt.Println(string(ByteMarshalerData))

	result := &dataSome{}
	json.Unmarshal(ByteMarshalerData, result)
	fmt.Println(*result)

	result2 := &dataSome{}
	json.Unmarshal(ByteMarshaler, result2)
	fmt.Println(*result2)
}

func case2() {
	src1 := map[string]interface{}{"num": 123}
	ByteMarshalerData, _ := jsonMarshaler.Marshal(src1)
	result1 := &data1{}
	json.Unmarshal(ByteMarshalerData, result1)
	fmt.Println(*result1)
}

func case3() {
	Result1 := &data{}
	jsonStr := `{"TestStringIntOtherName":0,"TestStringBool":true,"TTTT":true,"TestStringInt":"0","Name1":11,"Name2":22,"Name3":33}`
	json.Unmarshal([]byte(jsonStr), Result1)
	fmt.Println(*Result1)

	jsonStr = `{"TestStringInt":"0","Name1":11,"Name3":33,"TestStringIntOtherName":0,"TestStringBool":true,"name2":22,"TTTT":true}`
	json.Unmarshal([]byte(jsonStr), Result1)
	fmt.Println(*Result1)
}
