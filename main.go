package main

import (
	"os"
	"fmt"
)

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	defer app.Close()

	port := os.Getenv("PORT")
	fmt.Println("Listening on port " + port)
	app.Run(":" + port)
}
