package controllers

import (
	"fmt"
	"go-webapp/app/models"
	"log"
	"net/http"
)

// Topページに行った際のハンドラーを作る
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	// err == "invalid session"
	if err != nil {
		// ログインしていない場合は以下を行う
		generateHTML(w, "This is a top page", "layout", "public_navbar", "top")
	} else {
		// ログインしている場合は/todosに移動する
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

// /todosにアクセスがあった際のハンドラー
func todos(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		// sessionなし
		http.Redirect(w, r, "/top", http.StatusFound)
	} else {
		// sessionあり
		// 何をしている？
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// そのユーザーのtodoを全て取得する
		todos, _ := user.GetMultipleTodo()
		user.Todos = todos
		fmt.Println(user.Todos)
		// userをtodo_index.htmlに渡す -> .Todos, .ID と取得
		generateHTML(w, user, "layout", "private_navbar", "todo_index")
	}
}

// type User struct {
// 	ID        int
// 	UUID      string
// 	Name      string
// 	Email     string
// 	PassWord  string
// 	CreatedAt time.Time
// 	Todos     []Todo // [{1 First Todo 0 2022-06-29 16:25:14.146508 +0900 +0900} {2 second todo 0 2022-06-29 16:29:16.31486 +0900 +0900}]
// }

// type Todo struct {
// 	ID        int
// 	Content   string // 構造体の埋め込み
// 	UserID    int
// 	CreatedAt time.Time
// }

// todo create
func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	fmt.Println(err)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}

}

func todoSave(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// name属性の値を取得
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", http.StatusFound)

	}
}

// TODO: 確認必要
// 編集の際のハンドラーを作成する
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	fmt.Println("---")
	fmt.Println(id)
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		// ユーザーの取得
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		_, err = session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// 1件のtodoの取得
		todo, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, todo, "layout", "private_navbar", "todo_edit")

	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		todo := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := todo.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

// type Todo struct {
// 	ID        int
// 	Content   string
// 	UserID    int
// 	CreatedAt time.Time
// }
