package main

import (
    //"fmt"
    "log"
    "strconv"
    "net/http"
    "tasks/models"
    "tasks/auth"
    "tasks/utils"
    "tasks/rabbit"
    "tasks/database"
    "encoding/json"
    // "github.com/gin-gonic/gin"
    //"github.com/streadway/amqp"
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

func IdsTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
    var ids []string
    if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    message, err := json.Marshal(ids)
    if err != nil {
        http.Error(w, "Failed to serialize JSON", http.StatusInternalServerError)
        return
    }

    ch, _ := rabbit.ConnectRabbitMQ()
    err = rabbit.PublishMessage(ch, message)
    if err != nil {
        http.Error(w, "Couldn't send message to RabbitMQ", http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "IDs received and sent to RabbitMQ"}`))
}

func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/login", UserValidation)
    mux.HandleFunc("/task", CreateNewTask)
    mux.HandleFunc("/delete", deleteTask)
    mux.HandleFunc("/idtasks", IdsTask)

	err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)

}