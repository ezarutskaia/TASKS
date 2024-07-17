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
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    tocken, _ := auth.CreateSession(user.Email, user.Password)

	w.Write([]byte(tocken))
}

func SessionValidation(w http.ResponseWriter, r *http.Request) {
    db := *utils.Engine()
    var user models.User
    var requestData map[string]string
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    email := requestData["email"]
    uuid := requestData["uuid"]
    token := requestData["token"]
    taskname := requestData["taskname"]

    session := auth.IsSession(email, uuid, token)

    if session == true {
        result := db.Preload("Roles").Where("email = ?", email).First(&user)
        if result.Error != nil {
            log.Fatal(result.Error)
        }
        task, _ := user.CreateTask(taskname)
        result = db.Create(task)
        if result.Error != nil {
            log.Fatal(result.Error)
        }

        w.Write([]byte("session was created with id " + strconv.Itoa(task.Id)))
    } else {
        w.Write([]byte("session is not exist"))
    }
}

func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/login", UserValidation)
    mux.HandleFunc("/task", SessionValidation)

	err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)

}