package middlew

import (
	"net/http"

	"github.com/drc288/go_mux/database"
)

func CheckDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == 0 {
			http.Error(w, "Connection refuse with database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
