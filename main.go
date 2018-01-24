//main.go

package main

import "os"

//using env variables for database connection for now
func main() {
	application := API{}
	application.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	
	//ensureTableExists()

	application.Run(":8080")

}
