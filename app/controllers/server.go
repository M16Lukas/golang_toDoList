package controllers

import (
	"golang_toDoList/config"
	"net/http"
)

func StartMainServer() error {
	// css, jsファイルの呼び込み
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URL登録
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
