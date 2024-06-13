package main

import (
	"tasks/models"
	"tasks/utils"
)

func main() {
	DataBase := *utils.Engine()

	DataBase.AutoMigrate(&models.Role{})
	DataBase.AutoMigrate(&models.User{})
	DataBase.AutoMigrate(&models.Task{})
}