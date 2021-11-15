package main

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"strconv"

	//"github.com/jmoiron/sqlx"
	//"database/sql"
	"encoding/json"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
)

/* Добавить нового пользователя
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username": "user_1"}' \
  http://localhost:9000/users/add
*/

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username"`
	Created  int    `json:"-"`
}

func Add_users() {
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
		//db, err := sql.Open("postgres", "postgresql://db_user:db_user_pass@postgres(myapp_db)/app_db")
		//connStr := "user=db_user password=db_user_pass host=myapp_db dbname=app_db sslmode=disable"
		//db, err := sql.Open("postgres", connStr)
		//db, err := sqlx.Connect("postgres", "user=db_user password=db_user_pass host=myapp_db dbname=app_db sslmode=disable")
		//defer db.Close()
		//err = db.QueryRow("insert into user (username) values ($1) return index", "Skitsch").Scan(&id)
		//err = db.Ping()

		conn, err := pgx.Connect(context.Background(), "postgres://db_user:db_user_pass@myapp_db:5432/app_db")
		defer conn.Close(context.Background())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var id int = 0
		//err = conn.Ping(context.Background())
		//if _, err_exec := conn.Exec(context.Background(), "insert into my_user (username) values ($1)", "Skitsch");
		err = conn.QueryRow(context.Background(), "insert into my_user (username) values ($1) returning id", p.Username).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			//fmt.Fprintln(w, "Bad insert(")
			return
		} else {
			fmt.Fprintln(w, "Id созданного пльзователя", id)
		}
	})
}

/*------------------------------------------------------------------------*/

/* Создать новый чат между пользователями
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": ["<USER_ID_1>", "<USER_ID_2>"]}' \
  http://localhost:9000/chats/add
*/

func checkUsers(users []int, ctx context.Context, conn *pgx.Conn, w http.ResponseWriter) error {
	var err error
	var check bool = true
	for _, user := range users {
		err = conn.QueryRow(ctx, "select exists(select 1 from my_user where id = ($1))", user).Scan(&check)
		if check == false {
			//fmt.Fprintf(w, "User id = %v\n", user)
			err = errors.New("User with id = " + strconv.Itoa(user) + " does not exist")
			return err
		}
	}
	return nil
}

func fillChatUser(users []int, chat_id int, ctx context.Context, conn *pgx.Conn) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, user_id := range users {
		_, err = tx.Exec(ctx,"insert into chat_user (user_id, chat_id) values ($1, $2)", user_id, chat_id)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

type Chat struct {
	Id         int      `json:"-"`
	Name       string   `json:"name"`
	Users      []int 	`json:"users"`
	Created_at int      `json:"-"`
}

func Create_chat() {
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
		conn, err := pgx.Connect(context.Background(), "postgres://db_user:db_user_pass@myapp_db:5432/app_db")
		defer conn.Close(context.Background())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err = checkUsers(p.Users, context.Background(), conn, w); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		//return
		var id int
		err = conn.QueryRow(context.Background(), "insert into chat (name) values ($1) returning id", p.Name).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		} else {
			fmt.Fprintln(w, "Id созданного чата", id)
		}
		if err = fillChatUser(p.Users, id, context.Background(), conn); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
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
	Id         int    `json:"-"`
	Chat       string `json:"chat"`   // ссылка на идентификатор чата, в который было отправлено сообщение
	Author     string `json:"author"` // ссылка на идентификатор отправителя сообщения, отношение многие-к-одному
	Text       string `json:"text"`   // текст отправленного сообщения
	Created_at int    `json:"-"`      // время создания
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
