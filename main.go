package main

import (
	"fmt"
	"os"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	app := New(dbHost, dbName)
	defer app.Close()

	fmt.Println("Listening on port " + port)
	app.Run(":" + port)
}
