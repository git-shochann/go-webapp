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

// 複数のTodoを取得する
// '複数'だからスライスで返す
func GetMultipleTodo() (todos []Todo, err error) {
	GetMultipleTodoCmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(GetMultipleTodoCmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		// 複数を処理するけど脳内は1つを処理するとして考える。
		todo := Todo{}
		// 作成したtodoに埋め込む
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}

		// スライスに要素を追加する
		// []Todoなので[todo{}, todo{}, todo{}...]
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}