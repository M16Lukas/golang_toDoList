package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (user *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content, user_id, created_at
	) values (?, ?, ?)`

	_, err = Db.Exec(cmd, content, user.ID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select 
						id, content, user_id, created_at
					from
						todos
					where
						id = ?`

	todo = Todo{}
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select 
						id, content, user_id, created_at
					from
						todos`

	rows, err := Db.Query(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)

		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (user *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select 
						id, content, user_id, created_at
					from
						todos
					where
						user_id = ?`

	rows, err := Db.Query(cmd, user.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)

		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (todo *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id = ?
					where id = ?`

	_, err := Db.Exec(cmd, todo.Content, todo.UserID, todo.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (todo *Todo) DeleteTodo() error {
	cmd := `delete from todos where id = ?`
	_, err := Db.Exec(cmd, todo.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
