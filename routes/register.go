package routes

import (
	"encoding/json"
	"net/http"

	"github.com/drc288/go_mux/database"
	"github.com/drc288/go_mux/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Error validating the email", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password is necesari more than 6 characters", 400)
		return
	}

	_, findedEmail, _ := database.CheckEmail(t.Email)
	if findedEmail {
		http.Error(w, "User in Database", 400)
		return
	}

	_, status, err := database.InsertRegister(t)
	if err != nil {
		http.Error(w, "Error in the insertion of the user "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Error in the insertion of the user in the database ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
