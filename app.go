//app.go

package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)

//struct with references to the router and database to be used
type API struct {
	Router *mux.Router
	DB *sql.DB
}

//need method to initialize and run the API
//Initialize uses database information
func (application *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)
	
	var err error
	application.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	application.Router = mux.NewRouter()
}

//Run starts the API
func (a *App) Run(addr string) {}
