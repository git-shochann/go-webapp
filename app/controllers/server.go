package controllers

import (
	"fmt"
	"go-webapp/config"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	// スライスの作成
	var files []string

	// filenames == ["layout", "public_navbar", "signup"]
	for _, v := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", v))
	}
	// files == ["app/views/templates/layout.html", "app/views/templates/public_navbar.html", "app/views/templates/signup.html"]

	// パースしてキャッシュする?
	template := template.Must(template.ParseFiles(files...))

	// 実際にHTMLを表示する処理を行う dataは渡したいデータ?
	template.ExecuteTemplate(w, "layout", data)

}

func StartMainServer() error {

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// HandleFuncの引数にtop関数を渡す -> HandleFuncの第2引数と同じ型にする
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	return http.ListenAndServe(":"+config.Config.Port, nil)

}
