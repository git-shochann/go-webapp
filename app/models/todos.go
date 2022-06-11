package models

import (
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
	CreateTodoCmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`
	_, err = Db.Exec(CreateTodoCmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
