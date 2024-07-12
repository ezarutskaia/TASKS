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

func CreateSession(email string, password string) (*models.Session, error){
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
        fmt.Println(tocken)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        return session, nil
    } else {
        return &models.Session{}, errors.New("User haven't got this password")
    }
}

func IsSession (email string, uuid string, tocken string) bool {
    db := *utils.Engine()
	var session models.Session

    result := db.Where("email = ? AND uuid = ? AND endsession > NOW()", email, uuid).First(&session)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return false
        } else {
            log.Fatal(result.Error)
        }
    }

    ValidTocken, _ := GenerateJWT(email)
    if tocken == ValidTocken {
        return true
    }
    return false 
}
