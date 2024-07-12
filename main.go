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
    uuid := "79fb3589-5f11-4ce1-af8b-58be522b7290"
   tocken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Inh5ekBtYWlsLmNvbSJ9.Ee7O9cwe7wBFizHY1hvAN0wJBj9PH2m6MIGx6trsncQ"
    check := auth.IsSession(email, uuid, tocken)

    if check == true {
        var user models.User

        result := db.Preload("Roles").Where("email = ?", email).First(&user)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        task, err := user.CreateTask("TestNine")
        result = db.Create(task)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        fmt.Println(task, err)
    }

}
