package main

import (
	"fmt"
	"os"
)

func main() {
	var dbHost, dbName, port string
	var ok bool

	dbHost, ok = os.LookupEnv("DB_HOST")
	if !ok {
		dbHost = "localhost"
	}

	dbName, ok = os.LookupEnv("DB_NAME")
	if !ok {
		dbName = "schemed"
	}

	port, ok = os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	server := New(dbHost, dbName)
	defer server.Close()

	fmt.Println("Listening on " + dbHost + "@" + dbName + ":" + port)
	server.Run(":" + port)
}
