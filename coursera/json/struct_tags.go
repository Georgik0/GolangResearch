package main

import (
	"fmt"
	"encoding/json"
)

type User_tags struct {
	ID int `json:"user_id,string"`			// user_id - имя(переопределенное) из которого нужно распаковывать или в которое нужно запаковывать; string - тип запаковываемых данных
	Username string							// обработается как есть
	Address string `json:",omitempty"`		// первый аргумент пустой - значит имя остается каким было, omitempty - если поле пустое, то его не нужно писать в результат при запаковке
	Company string `json:"-"`				// "-" означает, что поле не нужно обрабатывать
}

func main() {
	 u := User_tags{
		 ID: 42,
		 Username: "Georgik",
		 Address: "",
		 Company: "School21",
	 }
	 result, err := json.Marshal(u)
	 if err != nil {
		 panic(err)
	 }
	 fmt.Printf("json string %s\n", result)
}