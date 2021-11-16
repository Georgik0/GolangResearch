package main

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"
)

func CheckUsers(users []int, ctx context.Context, conn *pgx.Conn, w http.ResponseWriter) error {
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

func CheckChats(chats []int, ctx context.Context, conn *pgx.Conn, w http.ResponseWriter) error {
	var err error
	var check bool = true
	for _, chat_id := range chats {
		err = conn.QueryRow(ctx, "select exists(select 1 from chat where id = ($1))", chat_id).Scan(&check)
		if check == false {
			//fmt.Fprintf(w, "User id = %v\n", user)
			err = errors.New("Chat with id = " + strconv.Itoa(chat_id) + " does not exist")
			return err
		}
	}
	return nil
}

func CheckMessage(user_id int, chat_id int, ctx context.Context, conn *pgx.Conn) error {
	var check bool = false
	err := conn.QueryRow(ctx, "select exists(select 1 from my_user where id = ($1))", user_id).Scan(&check)
	if check == false {
		err = errors.New("User with id = " + strconv.Itoa(user_id) + " does not exist")
		return err
	}
	err = conn.QueryRow(ctx, "select exists(select 1 from chat where id = ($1))", chat_id).Scan(&check)
	if check == false {
		err = errors.New("Chat with id = " + strconv.Itoa(user_id) + " does not exist")
		return err
	}
	return nil
}
