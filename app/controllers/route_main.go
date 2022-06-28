package controllers

import (
	"net/http"
)

// Topページに行った際のハンドラーを作る
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		// ログインしていない場合は以下を行う
		generateHTML(w, "Hello!", "layout", "private_navbar", "top")
	} else {
		// ログインしている場合は/todosに移動する
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

// indexのハンドラーの作成
func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/top", http.StatusFound)
	} else {
		generateHTML(w, "nil", "layout", "private_navbar", "index")
	}
}
