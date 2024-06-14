package main

import (
	"fmt"
	"tasks/models"
	"tasks/utils"
	"log"
)

func main() {
	db := *utils.Engine()

	var users []models.User
    result := db.Find(&users)
    if result.Error != nil {
        log.Fatal(result.Error)
    }

    // Обновление паролей для каждого пользователя
    for _, user := range users {
        // Генерация пароля (например, на основе имени пользователя)
        plainPassword := fmt.Sprintf("password_for_%s", user.Email)

        // Обновление пользователя с новым паролем
        user.Password = plainPassword
        db.Save(&user)
    }

    fmt.Println("Пароли успешно обновлены.")
}
