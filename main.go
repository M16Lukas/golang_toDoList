package main

import (
	"fmt"
	"golang_toDoList/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.LogFile)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.SQLDriver)

	// log.Println("test")

	// user := &models.User{}
	// user.Name = "test"
	// user.Email = "test@google.com"
	// user.PassWord = "testtest"
	// fmt.Println(user)
	// user.CreateUser()

	user, _ := models.GetUser(1)
	fmt.Println(user)

	user.Name = "test2"
	user.Email = "test2@google.com"
	user.UpdateUser()

	user, _ = models.GetUser(1)
	fmt.Println(user)
}
