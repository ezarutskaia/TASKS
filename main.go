package main

import (
	"fmt"
    "log"
    "tasks/utils"
    "tasks/auth"
    "tasks/models"
)

func main() {
    db := *utils.Engine()

	email := "ezv@mail.com"
    uuid := "a34fe3e5-781a-4d67-96e6-fff74ad4133d"
    check := auth.IsSession(email, uuid)

    if check == true {
        var user models.User

        result := db.Where("email = ?", email).First(&user)
        if result.Error != nil {
                log.Fatal(result.Error)
            }
        task, err := user.CreateTask("TestSix")
        result = db.Create(task)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        fmt.Println( task, err)
    }

}
