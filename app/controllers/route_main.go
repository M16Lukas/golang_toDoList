package controllers

import (
	"net/http"
	"text/template"
)

// ハンドラ
func top(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("app/views/templates/top.html")
	t.Execute(w, nil)
}