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
		_, err := session(w, r)
		if err != nil {
			// ログインしていない
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
		// GETの場合はそのままsignupのhtmlを表示する
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			// name属性から取得
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		fmt.Println("user created!")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// 実際のログインページを表示するハンドラー
func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		// ログインしていない
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

// ユーザーを認証するハンドラー
func authenticate(w http.ResponseWriter, r *http.Request) {
	// こうすることで後ほどrからPostリクエストの内容を取得出来る。
	// r.Form と r.PostFormに値をセットする。
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		// userが見つからなかったら
		// .logに書き込むだけ fatalln()はos.Exit(1)を呼び出してプログラムを終了させてしまう。
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)

	}
	// 複合化したpasswordと比較する
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		// セッションを作成する
		session, err := user.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}

		// クッキーを作成して、ブラウザに保存場所を決める
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}

		// レスポンスヘッダーにCookieをセットする
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}
}

// ログアウトのハンドラー
func logout(w http.ResponseWriter, r *http.Request) {

	// BUG: Cookieがない？ == 元々sessionが作成出来てない？
	// クッキーをまずはリクエストから取得する
	cookie, err := r.Cookie("_cookie")

	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		// session構造体を初期化してクッキーをセットする
		session := models.Session{UUID: cookie.Value}
		// 実際に削除を行う
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
