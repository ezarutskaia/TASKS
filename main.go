package main

import (
    "log"
    "net/http"
    "tasks/models"
    "tasks/auth"
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

func main() {

	mux := http.NewServeMux()
    mux.HandleFunc("/login", UserValidation)

	err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)

}