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

// 全部のTodoを取得する
// '全部'だからスライスで返す
func GetAllTodo() (todos []Todo, err error) {
	GetAllTodoCmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(GetAllTodoCmd)
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

// 特定ユーザーのTodoを全部取得する
func (u *User) GetMultipleTodo() (todos []Todo, err error) {
	GetMultipleTodoCmd := `select id, content, user_id, created_at from todos where user_id = ?`
	rows, err := Db.Query(GetMultipleTodoCmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// Todoの更新
func (t *Todo) UpdateTodo() (err error) {
	// todoのidを指定して、content_idとuser_idを変更する
	UpdateTodoCmd := `update todos set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(UpdateTodoCmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err

}

// Todoの削除
func (t *Todo) DeleteTodo() (err error) {
	DeleteTodoCmd := `delete from todos where id = ?`
	_, err = Db.Exec(DeleteTodoCmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
