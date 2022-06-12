package models

import (
	"database/sql"
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u User) CreateTodo(content string) (err error) {
	createTodoCmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`
	_, err = Db.Exec(createTodoCmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	getTodoCmd := `select id, content, user_id, created_at from todos where id = ?`
	todo = Todo{}
	row := Db.QueryRow(getTodoCmd, id)
	err = row.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("type is %T, value is %v\n", err, err) // ここの型とメッセージを調べてみた
		} else {
			log.Fatalln(err)
		}
	}

	return todo, err
}
