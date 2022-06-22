package controllers

import (
	"net/http"
)

// Topページに行った際のハンドラーを作る
func top(w http.ResponseWriter, r *http.Request) {

	// Helloを渡す？
	generateHTML(w, "Hello!", "layout", "public_navbar", "top")
}
