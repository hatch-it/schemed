
package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

// App contains all the state for the entire application.
type App struct {
	Router *gin.Engine
	DB *sql.DB
}

// Initialize takes the details required to connect to the database.
// Create a connection and wire up the routes to response accordingly.
func (a *App) Initialize(user, password, dbname string) {
	flags := fmt.Sprintf("host=localhost sslmode=disable user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", flags)
	if err != nil {
		log.Fatal(err)
	}

	provision(a.DB)

	a.Router = gin.Default()
	
	
}

// Run starts the application.
func (a *App) Run(addr string) {
	a.Router.Run()
}

// Provision the database with some basic models.
func provision(db *sql.DB) {
	log.Print("Provisioning the database (this might take a minute or two)")

	queries := []string{`
		CREATE EXTENSION IF NOT EXISTS pgcrypto; -- Enable gen_random_uuid()
		SET TIMEZONE TO 'Etc/UTC';               -- Set default timezone to UTC
	`, `
		-- User Table
		CREATE TABLE IF NOT EXISTS users(
			id         UUID         PRIMARY KEY     DEFAULT gen_random_uuid(),
			created_at TIMESTAMP    WITH TIME ZONE  DEFAULT now(),
			updated_at TIMESTAMP    WITH TIME ZONE  DEFAULT now(),
			deleted_at TIMESTAMP		WITH TIME ZONE,

			email      VARCHAR(254) NOT NULL,
			password   VARCHAR(74)  NOT NULL
		);
	`, `
		-- Venue Table
		CREATE TABLE IF NOT EXISTS venues(
			id         UUID          PRIMARY KEY,
			created_at TIMESTAMP     WITH TIME ZONE,
			updated_at TIMESTAMP     WITH TIME ZONE,
			deleted_at TIMESTAMP		 WITH TIME ZONE,

			google     VARCHAR(256), -- ID for Google API
			foursquare VARCHAR(256), -- ID for Foursquare API
			yelp       VARCHAR(256)  -- ID for Yelp API
		);
	`, `
		-- Event Table
		CREATE TABLE IF NOT EXISTS events(
			id         UUID         PRIMARY KEY,
			created_at TIMESTAMP    WITH TIME ZONE,
			updated_at TIMESTAMP    WITH TIME ZONE,
			deleted_at TIMESTAMP    WITH TIME ZONE,

			venue      UUID         REFERENCES venues,
			start_time TIMESTAMP    WITH TIME ZONE,
			end_time   TIMESTAMP    WITH TIME ZONE,

			facebook   VARCHAR(256) -- ID for Facebook API
		);
	`}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("Done!")
}