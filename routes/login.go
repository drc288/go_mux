package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/drc288/go_mux/database"
	"github.com/drc288/go_mux/jwt"
	"github.com/drc288/go_mux/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "User or password invalid"+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "The email has ben required", 400)
		return
	}

	userDocument, ok := database.LoginPass(u.Email, u.Password)
	if !ok {
		http.Error(w, "User or password invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(userDocument)
	if err != nil {
		http.Error(w, "Error generating the token"+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
