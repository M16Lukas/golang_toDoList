package main

import (
	"golang_toDoList/app/models"
)

func main() {

	// user := &models.User{}
	// user.Name = "test2"
	// user.Email = "test2@gmail.com"
	// user.PassWord = "testtest"
	// user.CreateUser()

	// user2, _ := models.GetUser(3)
	// user2.CreateTodo("Third todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// ts, _ := user2.GetTodosByUser()
	// for _, v := range ts {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(1)
	t.Content = "Update Todo"
	t.UpdateTodo()
	// fmt.Println(t)
}
