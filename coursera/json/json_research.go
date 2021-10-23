package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	ID int
	Username string
	phone string
}

var jsonStr = `{"id": 42, "username": "rvasily", "phone": "123"}`

func main() {
	data := []byte(jsonStr)

	pu := &User{}
	u := User{}
	fmt.Printf("data = %v   %[1]T\nuser = %v   %[2]T\npuser = %v   %[3]T\n", data, u, pu)

	json.Unmarshal(data, pu)
	fmt.Printf("pu = %v   *pu = %v\n", pu, *pu)
	pu.Username = "newName"

	result, err := json.Marshal(pu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json = %s  %[1]T\n", string(result))
}
