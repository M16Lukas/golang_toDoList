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
	Todos     []Todo
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
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

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
					from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
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

func (user *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions(
		uuid, email, user_id, created_at
	) values (?, ?, ?, ?)`
	_, err = Db.Exec(cmd1, createUUID(), user.Email, user.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at
						from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(cmd2, user.ID, user.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id ,uuid, email, user_id, created_at
					from sessions where uuid = ?`
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt,
	)

	if err != nil {
		valid = false
		return
	}

	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at
					from users
					where id = ?`

	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	return user, err
}
