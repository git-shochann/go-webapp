package controllers

import (
	"fmt"
	"go-webapp/app/models"
	"log"
	"net/http"
)

// 新規登録のハンドラーを作成 == 新規登録ページのパスにアクセスがあったらこの関数を渡す
func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// GETの場合はそのままsignupのhtmlを表示する
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
		fmt.Println("user created!")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}

// ユーザーの認証のハンドラー
func authenticate(w http.ResponseWriter, r *http.Request) {
	// こうすることで後ほどrからPostリクエストの内容を取得出来る。
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		// userが見つからなかったら
		// .logに書き込むだけ fatalln()はos.Exit(1)を呼び出してプログラムを終了させてしまう。
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)

	}
	// 複合化したpasswordと比較する
	if user.PassWord == models.Encrypto(r.PostFormValue("password")) {
		// セッションを作成する
		session, err := user.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}

		// クッキーを作成して、ブラウザに保存場所を決める
		cookie := http.Cookie{
			Name:     "_value",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/top", http.StatusFound)
	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}
}
