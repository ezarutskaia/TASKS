package main

import (
	"fmt"
	"tasks/models"
)

func main() {
	// two := &models.User{Email: "oen@mail.com", Id: 2}

	// two.AddRole(models.Role{"Worker", 3})
	// two.AddRole(models.Role{"Admin", 7})
	// two.AddRole(models.Role{"Glavbuh", 5})
	// two.RevokeRole(models.Role{"Glavbuh", 5})

	// oleg := &models.Manager{User: *two, AccesLevel: 3}
	// zadacha, _ := oleg.CreateTask("new")

	// fmt.Println(*oleg, zadacha)
	// fmt.Println(two)

	tasks := []models.Task{
		models.Task{Id: 1, Name: "A", UserId: 2, Status: models.InProgress, Priority: models.Normal},
		models.Task{Id: 2, Name: "B", UserId: 2, Status: models.InProgress, Priority: models.Normal},
		models.Task{Id: 3, Name: "C", UserId: 2, Status: models.InProgress, Priority: models.Low},
		models.Task{Id: 4, Name: "D", UserId: 2, Status: models.InProgress, Priority: models.High},
		models.Task{Id: 5, Name: "F", UserId: 2, Status: models.Open, Priority: models.Critical},
	}
	filters := []models.TaskFilter{
		models.TaskFilter{Type: "status", Value: models.Open},
		models.TaskFilter{Type: "priority", Value: models.Critical},
	}
	filtertasks := models.FilterTaskSlice(tasks, filters)

	fmt.Println(filtertasks)
}
