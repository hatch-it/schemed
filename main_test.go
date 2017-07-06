package main_test

import (
	"os"
	"log"
	"testing"

	"github.com/puradox/schemed"
)

var app main.App

func TestMain(m *testing.M) {
	app = main.App{}
	app.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
	)

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `
	-- User Table
	CREATE TABLE IF NOT EXISTS users(
		id         UUID         PRIMARY KEY     DEFAULT gen_random_uuid(),
		created_at TIMESTAMP    WITH TIME ZONE  DEFAULT now(),
		updated_at TIMESTAMP    WITH TIME ZONE  DEFAULT now(),
		deleted_at TIMESTAMP		WITH TIME ZONE,

		email      VARCHAR(254) NOT NULL,
		password   VARCHAR(74)  NOT NULL
	);
`

func clearTable() {
	app.DB.Exec("DELETE FROM users")
}