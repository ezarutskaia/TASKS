package utils

import (
	"gorm.io/gorm"
    "gorm.io/driver/mysql"
	"log"
)

func Engine() **gorm.DB {
	DSN := "root:secret@tcp(0.0.0.0:4306)/tasks?parseTime=true"
	DataBase, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
        log.Fatal(err)
    }
	return &DataBase
}