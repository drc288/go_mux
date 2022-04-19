package routes

import (
	"encoding/json"
	"net/http"

	"github.com/drc288/go_mux/database"
	"github.com/drc288/go_mux/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Data are invalid"+err.Error(), 400)
		return
	}

	var status bool

	status, err = database.ModifyRegister(u, IDUsuario)
	if err != nil {
		http.Error(w, "Error trying to modify the profile", 400)
		return
	}

	if !status {
		http.Error(w, "Unable to modificate the profile", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
