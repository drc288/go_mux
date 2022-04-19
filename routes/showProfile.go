package routes

import (
	"encoding/json"
	"net/http"

	"github.com/drc288/go_mux/database"
)

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter id", http.StatusBadRequest)
		return
	}

	profile, err := database.FindProfile(ID)
	if err != nil {
		http.Error(w, "Error searching the register "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(profile)
}
