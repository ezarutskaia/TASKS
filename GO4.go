package main

import (
	"fmt"
	"tasks/models"
)

func main() {
	two := &models.User{Email: "oen@mail.com", Id: 2}

	two.AddRole(models.Role{"Worker", 3})
	two.AddRole(models.Role{"Admin", 7})
	two.AddRole(models.Role{"Glavbuh", 5})
	two.RevokeRole(models.Role{"Glavbuh", 5})

	oleg := &models.Manager{User: *two, AccesLevel: 3}
	zadacha, _ := oleg.CreateTask("new")

	fmt.Println(*oleg, zadacha)
	fmt.Println(two)
}
