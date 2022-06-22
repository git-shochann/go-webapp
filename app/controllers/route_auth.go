package controllers

import "net/http"

// 新規登録のハンドラーを作成
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "signup")
}
