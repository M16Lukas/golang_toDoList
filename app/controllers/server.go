package controllers

import (
	"golang_toDoList/config"
	"net/http"
)

func StartMainServer() error {
	// URL登録
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
