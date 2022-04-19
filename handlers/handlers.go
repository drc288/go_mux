package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/drc288/go_mux/middlew"
	"github.com/drc288/go_mux/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Mannage routes, CORS & Run the sercice
func Mannage() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDatabase(routes.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDatabase(routes.Login)).Methods("POST")
	router.HandleFunc("/showprofile", middlew.CheckDatabase(middlew.ValidateJWT(routes.ShowProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middlew.CheckDatabase(middlew.ValidateJWT(routes.ModifyProfile))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Printf("using: http://localhost:%s/\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
