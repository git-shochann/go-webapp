package controllers

import (
	"net/http"
)

// ハンドラーを作る
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "layout", "top")
}
