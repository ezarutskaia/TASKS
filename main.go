package main

import (
	"fmt"
	"tasks/models"
	"tasks/utils"
)

func main() {
	db := *utils.Engine()

	// role := []models.Role{
	// {Name: "Developer", Value: 6}, {Name: "Tester", Value: 5}, 
	// {Name: "Glavbuh", Value: 4}, {Name: "Worker", Value: 3}, {Name: "Admin", Value: 7},}

	// data := []models.User{
	// 	{Email: "xyz@mail.com", Roles: []models.Role{models.Role{Id: 4, Name:"Worker", Value: 3},}},
	// 	{Email: "ezv@mail.com", Roles: []models.Role{models.Role{Id: 1, Name:"Developer", Value: 6}, models.Role{Id: 2, Name:"Tester", Value: 5},}},
	// 	{Email: "oen@mail.com", Roles: []models.Role{models.Role{Id: 5, Name:"Admin", Value: 7}, models.Role{Id: 3, Name:"Glavbuh", Value: 4},}},
	// 	{Email: "44@mail.com", Roles: []models.Role{models.Role{Id: 4, Name:"Worker", Value: 3}, models.Role{Id: 2, Name:"Tester", Value: 5},}},
	// }

	data := []models.Task{
		{Name: "testOne", UserID: 3, Status: models.InProgress, Priority: models.Normal},
		{Name: "testTwo", UserID: 2, Status: models.Open, Priority: models.Low},
		{Name: "testThree", UserID: 1, Status: models.Closed, Priority: models.Critical},
		{Name: "testFour", UserID: 3, Status: models.InProgress, Priority: models.Normal},
		{Name: "testFive", UserID: 4, Status: models.Done, Priority: models.High},
	}

	 result := db.Create(&data)
	 if result.Error != nil {
	 	fmt.Println(result.Error)
	 }
	var watch []models.Task
	db.Find(&watch)

	fmt.Println(watch)
}
