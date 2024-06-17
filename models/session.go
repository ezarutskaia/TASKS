package models

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Id    int `gorm:"primaryKey"`
	Email string 
	Uuid string
}