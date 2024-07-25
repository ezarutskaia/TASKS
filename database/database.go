package database

import (
	"fmt"
	"log"
	"time"
	"errors"
	"tasks/models"
	"gorm.io/gorm"
    "github.com/google/uuid"
)

func GetUser (db *gorm.DB, email string) (user *models.User) {
	result := db.Preload("Roles").Where("email = ?", email).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            fmt.Printf("User with email %s not found.\n", email)
        } else {
            log.Fatal(result.Error)
        }
    }
	return user
}

func CreateSession (db *gorm.DB, email string) {
	now := time.Now()
	session := &models.Session{Email: email, Uuid: uuid.New().String(), Endsession: now.Add(time.Hour)}
	result := db.Create(session)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func GetSession (db *gorm.DB, email string) (session *models.Session, err error) {
    result := db.Where("email = ? AND endsession > NOW()", email).Last(&session)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
			return &models.Session{}, errors.New("Active session not found.")
        } else {
            log.Fatal(result.Error)
        }
    }
	return session, nil
}

func CreateTask (db *gorm.DB, taskname string, user *models.User) (task *models.Task){
	task, _ = user.CreateTask(taskname)
	result := db.Create(task)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return task
}

func DeleteNoteByID (db *gorm.DB, model interface{}, id string) {
	result := db.Delete(model, id)

	if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            fmt.Printf("Note with id %s not found.\n", id)
        } else {
            log.Fatal(result.Error)
        }
	}
}