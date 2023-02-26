package main

import (
	"fmt"
	"golang_toDoList/app/models"
)

func main() {

	// user := &models.User{}
	// user.Name = "test"
	// user.Email = "test@gmail.com"
	// user.PassWord = "testtest"
	// user.CreateUser()

	// user2, _ := models.GetUser(2)
	// user2.CreateTodo("first todo")

	t, _ := models.GetTodo(1)
	fmt.Println(t)
}
