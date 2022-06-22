package controllers

import (
	"fmt"
	"go-webapp/config"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string               // ["A", "B", "C"]
	for _, file := range filenames { // スライスを回す
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// 実際の作成されたスライスの中身を確認する
	fmt.Println(files)

	// パースしてキャッシュする？
	template := template.Must(template.ParseFiles(files...))
	// 第二引数は実行するテンプレート、第3引数は
	template.ExecuteTemplate(w, "layout", data)

}

func StartMainServer() error {

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// HandleFuncの引数にtop関数を渡す -> HandleFuncの第2引数と同じ型にする
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	return http.ListenAndServe(":"+config.Config.Port, nil)

}
