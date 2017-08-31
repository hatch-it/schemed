package main

import (
	"fmt"
	"os"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbName := getEnv("DB_NAME", "schemed")
	port := getEnv("PORT", "3000")

	server := New(dbHost, dbName)
	defer server.Close()

	fmt.Println("Listening on " + dbHost + "@" + dbName + ":" + port)
	server.Run(":" + port)
}
