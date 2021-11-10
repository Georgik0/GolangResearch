package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/* Добавить нового пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username": "user_1"}' \
  http://localhost:9000/users/add
*/

type User struct {
	Id int	`json:"-"`
	Username string `json:"username"`
	Created int `json:"-"`
}

func Add_users() {
	id := 0 // где-то сохранить id последнего пользователя (чтобы не терялся при перезапуске программы)
	http.HandleFunc("/users/add", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		p := &User{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		/* Здесь нужно добавить пользователя */
		db, err := sql.Open("mysql", "db_user:db_user_pass@tcp(myapp_db)/app_db")
		defer db.Close()
		err = db.Ping()
		if err != nil {
			fmt.Fprintln(w, "No connect with mysql")
			return
		} else {
			fmt.Fprintln(w,"Get connect mysql")
		}

		fmt.Fprintf(w, "id созданного пользователя: %v\n", id)
		id++
	})
}
/*------------------------------------------------------------------------*/

/* Создать новый чат между пользователями
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": ["<USER_ID_1>", "<USER_ID_2>"]}' \
  http://localhost:9000/chats/add
*/

type Chat struct {
	Id int `json:"-"`
	Name string `json:"name"`
	Users []string `json:"users"`
	Created_at int `json:"-"`
}

func Create_chat() {
	id := 0
	http.HandleFunc("/chats/add", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		p := &Chat{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintln(w, "Chat: ", p)
		/* Здесь нужно добавить чат в бд */
		fmt.Fprintf(w, "id созданного чата: %v\n", id)
		id++
	})
}
/*------------------------------------------------------------------------*/

/* Отправить сообщение в чат от лица пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": "<CHAT_ID>", "author": "<USER_ID>", "text": "hi"}' \
  http://localhost:9000/messages/add
*/

type Message struct {
	Id int `json:"-"`
	Chat string `json:"chat"` // ссылка на идентификатор чата, в который было отправлено сообщение
	Author string `json:"author"` // ссылка на идентификатор отправителя сообщения, отношение многие-к-одному
	Text string `json:"text"` // текст отправленного сообщения
	Created_at int `json:"-"` // время создания
}

func Message_add() {
	id := 0
	http.HandleFunc("/messages/add", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		//fmt.Println(string(body))
		p := &Message{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintln(w, "Message: ", p)
		/* Здесь нужно добавить сообщение в бд */
		fmt.Fprintf(w, "id отправленного сообщения: %v\n", id)
		id++
	})
}
/*------------------------------------------------------------------------*/

/* Получить список чатов конкретного пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user": "<USER_ID>"}' \
  http://localhost:9000/chats/get
*/

type User_chats struct {
	User string `json:"user"`
}

func Get_slice_chats() {
	http.HandleFunc("/chats/get", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		//fmt.Println(string(body))
		p := &User_chats{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintln(w, "User: ", p)
		/* Здесь нужно получить список чатов этого пользователя из бд */
	})
}
/*------------------------------------------------------------------------*/

/* Получить список сообщений в конкретном чате
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": "<CHAT_ID>"}' \
  http://localhost:9000/messages/get
*/

type Chat_messages struct {
	Chat string `json:"chat"`
}

func Get_slice_messages() {
	http.HandleFunc("/messages/get", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		//fmt.Println(string(body))
		p := &Chat_messages{}
		err = json.Unmarshal(body, p)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintln(w, "Chat: ", p)
		/* Здесь нужно получить список сообщений этого чата из бд */
	})
}

func main() {
	Add_users()
	Create_chat()
	Message_add()
	Get_slice_chats()
	Get_slice_messages()
	http.ListenAndServe(":9000", nil)
}
