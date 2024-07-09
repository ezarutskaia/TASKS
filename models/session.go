package models

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	Id    int `gorm:"primaryKey"`
	Email string 
	Uuid string
	Endsession time.Time
}