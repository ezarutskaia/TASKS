package main

import (
	"fmt"
    "log"
    "tasks/utils"
    "tasks/models"
    "tasks/auth"
)

// func main() {

// 	email := "xyz@mail.com"
//     password := "password_for_xyz@mail.com"

//     auth.CreateSession(email, password)

// }

func main() {
    db := *utils.Engine()

    email := "xyz@mail.com"
    uuid := "a9bd74b3-c427-4291-a6e9-b09cb986d43c"
    check := auth.IsSession(email, uuid)

    if check == true {
        var user models.User

        result := db.Preload("Roles").Where("email = ?", email).First(&user)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        task, err := user.CreateTask("TestSeven")
        result = db.Create(task)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        fmt.Println(task, err)
    }

}
