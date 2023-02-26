package main

import (
	"fmt"
	"golang_toDoList/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.LogFile)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.SQLDriver)
}
