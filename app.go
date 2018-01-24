//app.go

package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//struct with references to the router and database to be used
type API struct {
	Router *mux.Router
	DB *sql.DB
}

//need method to initialize and run the API
//Initialize uses database information
func (a *App) Initialize(user, password, dbname string) {}

//Run starts the API
func (a *App) Run(addr string) {}
