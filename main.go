package main

import (
    "log"
    "strconv"
    "net/http"
    "tasks/models"
    "tasks/auth"
    "tasks/utils"
    "tasks/database"
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
        user := database.GetUser(db, emailHeader)
        task := database.CreateTask(db, taskname, user)

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
        database.DeleteNoteByID(db, &models.Task{}, id)
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