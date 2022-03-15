package api

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     sql.DB
}

func (a *App) Init(user string, password string, dbname string) {

}

func (a *App) Run(user string, password string, dbname string) {

}
