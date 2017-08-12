package main

import (
	"os"
	"fmt"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	app := App{}
	app.Initialize(dbHost, dbName)
	defer app.Close()

	fmt.Println("Listening on port " + port)
	app.Run(":" + port)
}
