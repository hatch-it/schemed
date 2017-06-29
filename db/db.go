package db

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
)

type DB sql.DB

// Open a connection to the database.
func Open() DB {
	var production bool
	var dbName, dbUser, dbPassword string

	flag.BoolVar(&production, "production", false, "optimize execution for production")
	flag.StringVar(&dbName, "db", "schemed", "database name")
	flag.StringVar(&dbUser, "user", "app", "database username")
	flag.StringVar(&dbPassword, "password", "", "database password")
	flag.Parse()

	var options string

	if production {
		options = "host=localhost dbname=" + dbName + " user=" + dbUser
	} else {
		options = "host=localhost sslmode=disable dbname=" + dbName + " user=" + dbUser + " password=" + dbPassword
	}

	db, err := sql.Open("postgres", options)
	if err != nil {
		log.Fatal(err)
	}

	setup(db)

	return db
}

// Run an insertion into an existing table with the given fields.
// Returns the generated id.
func Insert(db *sql.DB, table string, fields map[string]interface{}) string {
	queries := []string{
		"INSERT INTO ",
		table,
		"(",
		joinKeys(fields, ", "),
		") ",
		"VALUES (",
		joinValues(fields, ", "),
		") RETURNING id;",
	}

	var id string
	err := db.QueryRow(strings.join(query)).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	return id
}

// Run some preliminary queries after connecting to the database.
func setup(db *sql.DB) {
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
