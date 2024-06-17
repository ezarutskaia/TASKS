package auth

import (
    "log"
    "fmt"
    "errors"
	"tasks/models"
	"tasks/utils"
    "gorm.io/gorm"
    "github.com/google/uuid"
)

func CreateSession(e string, p string) (*models.Session, error){
    db := *utils.Engine()
	var user models.User

    result := db.Where("email = ?", e).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            fmt.Printf("Пользователь с email %s не найден.\n", e)
        } else {
            log.Fatal(result.Error)
        }
    } 

    if user.Password == p {
        session := &models.Session{Email: e, Uuid: uuid.New().String()}
        result := db.Create(session)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        return session, nil
    } else {
        return &models.Session{}, errors.New("User haven't got this password")
    }
}

func IsSession (e string, u string) bool {
    db := *utils.Engine()
	var session models.Session

    result := db.Where("email = ? AND uuid = ?", e, u).First(&session)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return false
        } else {
            log.Fatal(result.Error)
        }
    }
    return true 
}
