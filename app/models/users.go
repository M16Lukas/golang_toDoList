package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

// User 生成
func (user *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		user.Name,
		user.Email,
		Encrypt(user.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// User情報取得
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
					from users
					where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

// Userの更新
func (user *User) UpdateUser() (err error) {
	cmd := `update users 
					set name = ?, email = ?
					where id = ?`

	_, err = Db.Exec(cmd, user.Name, user.Email, user.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Userの削除
func (user *User) DeleteUser() (err error) {
	cmd := `delete from users
					where id = ?`

	_, err = Db.Exec(cmd, user.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
