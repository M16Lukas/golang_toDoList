package models

import (
	"database/sql"
	"fmt"
	"golang_toDoList/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

// テーブル名宣言
const (
	tableNameUser = "users"
)

// テーブル作成
func init() {
	Db, err := sql.Open(
		config.Config.SQLDriver,
		config.Config.DbName,
	)

	// エラーハンドリング
	if err != nil {
		log.Fatal(err)
	}

	// コマンド作成
	cmdCreateUserTable := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s(
			id 					INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid 				STRING NOT NULL UNIQUE,
			name 				STRING,
			email 			STRING,
			password 		STRING,
			created_at 	DATETIME
		)`, tableNameUser)

	Db.Exec(cmdCreateUserTable)
}
