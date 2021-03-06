//model.go
package main

import (
	"database/sql"
	"errors"
)

//struct to represent example class for testing routes
type class struct {
	ID int `json."id"`
	Name string `json:"name"`
	Professor string `json:"professor"`
}

//define functions that will work on this class struct
func (p *class) createClass(db *sql.DB) error {
	return errors.New("createClass method not implemented")
}

func (p *class) deleteClass(db *sql.DB) error {
	return errors.New("deleteClass method not implemented")
}

func (p *class) udpateClass(db *sql.DB) error {
	return errors.New("updateClass method not implemented")
}

func (p *class) getClass(db *sql.DB) error {
	return errors.New("getClass method not implemented")
}

//define function that will fetch all classes
func getClasses(db *sql.DB, start, count int) ([]class, error) {
	return nil, errors.New("GetClasses methos not implemented")
}
