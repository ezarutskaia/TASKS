package main

import (
    "log"
    "strconv"
    "net/http"
    "tasks/models"
    "tasks/auth"
    "tasks/utils"
    "encoding/json"
)

func UserValidation(w http.ResponseWriter, r *http.Request) {
    emailHeader := r.Header.Get("email")
	passwordHeader := r.Header.Get("password")

    if emailHeader == "" || passwordHeader == "" {
		http.Error(w, "email or password not found", http.StatusBadRequest)
		return
	}

    token, _ := auth.CreateSession(emailHeader, passwordHeader)

	w.Write([]byte(token))
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
    db := *utils.Engine()
    var user models.User
    var requestData map[string]string
    emailHeader := r.Header.Get("email")
	tokenHeader := r.Header.Get("token")

    if emailHeader == "" || tokenHeader == "" {
		http.Error(w, "email or token not found", http.StatusBadRequest)
		return
	}

    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    taskname := requestData["taskname"]

    session := auth.IsSession(emailHeader, tokenHeader)

    if session == true {
        result := db.Preload("Roles").Where("email = ?", emailHeader).First(&user)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        task, _ := user.CreateTask(taskname)
        result = db.Create(task)
        if result.Error != nil {
            log.Fatal(result.Error)
        }

        w.Write([]byte("task was created with id " + strconv.Itoa(task.Id)))
    } else {
        w.Write([]byte("session is not exist"))
    }
}

func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/login", UserValidation)
    mux.HandleFunc("/task", CreateNewTask)

	err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)

}