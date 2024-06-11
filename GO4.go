package main

import (
//	"fmt"
	"tasks/models"
	"gorm.io/gorm"
    "gorm.io/driver/mysql"
	"log"
)

func main() {

	dsn := "root:secret@tcp(0.0.0.0:4306)/tasks"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Role{})
	role := models.Role{Name: "Admin", Value: 7}
	result := db.Create(&role)
	if result.Error != nil {
        log.Fatal(err)
    }

}
