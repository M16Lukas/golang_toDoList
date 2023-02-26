package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"golang_toDoList/config"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

// テーブル名宣言
const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

// テーブル作成
func init() {
	Db, err = sql.Open(
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

	cmdCreateTodoTable := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s(
			id 					INTEGER PRIMARY KEY AUTOINCREMENT,
			content 		TEXT,
			user_id 		INTEGER,
			created_at 	DATETIME
		)`, tableNameTodo)

	Db.Exec(cmdCreateTodoTable)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%s", sha1.Sum([]byte(plaintext)))
	return crypttext
}
