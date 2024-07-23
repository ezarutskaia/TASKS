package auth

import (
    "errors"
    "tasks/utils"
    "tasks/database"
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

func GetTokenSession(email string, password string) (string, error){
    db := *utils.Engine()
    user := database.GetUser(db, email)

    if user.Password == password {
        database.CreateSession(db, email)
        tocken, _ := GenerateJWT(email)
        return tocken, nil
    } else {
        return "", errors.New("User haven't got this password")
    }
}

func IsSession (email string, token string) bool {
    db := *utils.Engine()
    _, err := database.GetSession(db, email)

    if err == nil {
        ValidToken, _ := GenerateJWT(email)
        if token == ValidToken {
            return true
        }
    } 
    return false
}