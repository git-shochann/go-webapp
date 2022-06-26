package controllers

import (
	"go-webapp/app/models"
	"log"
	"net/http"
	"time"
)

// Topページに行った際のハンドラーを作る
func top(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Helloを渡す？
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			// name属性から取得
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Panicln(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}
