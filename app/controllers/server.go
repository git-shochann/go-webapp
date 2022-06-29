package controllers

import (
	"errors"
	"fmt"
	"go-webapp/app/models"
	"go-webapp/config"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	// スライスの作成
	var files []string

	// filenames == ["layout", "public_navbar", "signup"]
	for _, v := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", v))
	}
	// files == ["app/views/templates/layout.html", "app/views/templates/public_navbar.html", "app/views/templates/signup.html"]

	// パースしてキャッシュする
	template := template.Must(template.ParseFiles(files...))

	// TODO: 再度確認
	template.ExecuteTemplate(w, "layout", data)

}

// TODO: 再度確認
func session(w http.ResponseWriter, r *http.Request) (session models.Session, err error) {
	// リクエストのクッキーを取得する
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		// uuidがDBにあるか存在するか確認する
		session := models.Session{UUID: cookie.Value}

		// 存在しない場合だけエラーを返す
		if ok, _ := session.CheckSession(); !ok {
			err = errors.New("invalid session")
		}
	}
	// 存在している場合だとerrはnilになる
	return session, err
}

// todos/edit/1
var validPath = regexp.MustCompile("^/todos/(edit/update)/([0-9]+$)")

// 関数を引数に取る(func todoEdit)
// TODO: 再度確認必要
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todos/edit/1
		// 一致する部分をスライスで取り出す
		fmt.Println(r.URL.Path) //

		q := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Println(q) // []
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		fmt.Println(qi)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func StartMainServer() error {

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// HandleFuncの引数にtop関数を渡す -> HandleFuncの第2引数と同じ型にする
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", todos)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	return http.ListenAndServe(":"+config.Config.Port, nil)

}
