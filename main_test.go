//main_test.go

package main_test

import (
	"os"
	"testing"

	"."
)


const tableCreationQuery = `CREATE TABLE IF NOT EXISTS classes
(
id SERIAL
name TEXT NOT NULL,
professor TEXT NOT NULL,
CONSTRAINT classes_pkey PRIMARY KEY (id)
)`

//represents the application we are going to test
var application main.App

func TestMain(m *testing.M) {
	application = main.App{}
	application.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	createTestTable()

	code := m.Run()

	cleanDatabase()

	os.Exit(code)
}

func createTestTable() {
	if _, err := application.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func cleanDatabase() {
	application.DB.Exec("DELETE FROM classes")
	application.DB.Exec("ALTER SEQUENCE classes_id_seq RESTART WITH 1")
}
