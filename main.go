//main.go

package main

import "os"

//using env variables for database connection for now
func main() {
	application := App{}
	application.Initialize(
		os.Getenv("TEST_DB_USERNAME")
		os.Getenv("TEST_DB_PASSWORD")
		os.Getenv("TEST_DB_NAME")
	)
	
	//ensureTableExists()

	application.Run(":8080")

}
