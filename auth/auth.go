package auth

import (
    "log"
    "fmt"
    "time"
    "errors"
	"tasks/models"
	"tasks/utils"
    "gorm.io/gorm"
    "github.com/google/uuid"
    "github.com/golang-jwt/jwt"
)

var jwtKey = []byte("ordnung")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string) (string, error) {
	claims := &Claims{
		Email: email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateSession(email string, password string) (string, error){
    db := *utils.Engine()
	var user models.User
    now := time.Now()

    result := db.Where("email = ?", email).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            fmt.Printf("Пользователь с email %s не найден.\n", email)
        } else {
            log.Fatal(result.Error)
        }
    } 

    if user.Password == password {
        session := &models.Session{Email: email, Uuid: uuid.New().String(), Endsession: now.Add(time.Hour)}
        result := db.Create(session)
        tocken, _ := GenerateJWT(email)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        return tocken, nil
    } else {
        return "", errors.New("User haven't got this password")
    }
}

func IsSession (email string, token string) bool {
    db := *utils.Engine()
	var session models.Session

    result := db.Where("email = ? AND endsession > NOW()", email).Last(&session)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return false
        } else {
            log.Fatal(result.Error)
        }
    }

    ValidToken, _ := GenerateJWT(email)
    if token == ValidToken {
        return true
    }
    return false 
}
