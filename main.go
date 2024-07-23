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
    if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

    emailHeader := r.Header.Get("email")
	passwordHeader := r.Header.Get("password")

    if emailHeader == "" || passwordHeader == "" {
		http.Error(w, "email or password not found", http.StatusBadRequest)
		return
	}

    token, _ := auth.GetTokenSession(emailHeader, passwordHeader)

	w.Write([]byte(token))
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

func deleteTask(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

    db := *utils.Engine()

    queryParams := r.URL.Query()
	id := queryParams.Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

    emailHeader := r.Header.Get("email")
	tokenHeader := r.Header.Get("token")

    if emailHeader == "" || tokenHeader == "" {
		http.Error(w, "email or token not found", http.StatusBadRequest)
		return
	}

    session := auth.IsSession(emailHeader, tokenHeader)

    if session == true {
        result := db.Delete(&models.Task{}, id)

        if result.Error != nil {
            http.Error(w, "Database error: " + result.Error.Error(), http.StatusInternalServerError)
            return
        }

        if result.RowsAffected == 0 {
            http.Error(w, "No record found", http.StatusNotFound)
            return
        }

        w.Write([]byte("task was delete with id " + id))
    } else {
        w.Write([]byte("session is not exist"))
    }
}

func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/login", UserValidation)
    mux.HandleFunc("/task", CreateNewTask)
    mux.HandleFunc("/delete", deleteTask)

	err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)

}