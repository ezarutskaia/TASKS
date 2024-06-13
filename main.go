package main

import (
	"fmt"
	"tasks/models"
	"gorm.io/gorm"
    "gorm.io/driver/mysql"
	"log"
)

func main() {

	// role := []models.Role{
	// 	{Name: "Developer", Value: 6}, {Name: "Tester", Value: 5}, 
	// 	{Name: "Glavbuh", Value: 4}, {Name: "Worker", Value: 3},}
	// result := db.Create(&role)
	var roles []models.Role
	result := db.Find(&roles)
	if result.Error != nil {
        log.Fatal(err)
    }

	fmt.Println(roles)
}
