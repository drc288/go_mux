package middlew

import (
	"net/http"

	"github.com/drc288/go_mux/routes"
)

// ValidateJWT valudate the jwt in the request body
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error in the token !"+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
