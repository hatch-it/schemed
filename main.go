package main

import (
  "log"
  "database/sql"
  _ "github.com/lib/pq"
)

func setup(db *sql.DB) {
  log.Print("Provisioning the database (this might take a minute or two)")

  queries := []string{`
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- Enable uuid_generate_v4()
    SET TIMEZONE TO 'Etc/UTC';                  -- Set default timezone to UTC
  `,`
    -- User Table
    CREATE TABLE IF NOT EXISTS users(
      id       UUID         PRIMARY KEY,
      email    VARCHAR(254) NOT NULL,
      password VARCHAR(74)  NOT NULL
    );
  `,`
    -- Venue Table
    CREATE TABLE IF NOT EXISTS venues(
      id         UUID          PRIMARY KEY,
      google     VARCHAR(256),              -- ID for Google API
      foursquare VARCHAR(256),              -- ID for Foursquare API
      yelp       VARCHAR(256)               -- ID for Yelp API
    );
  `,`
    -- Event Table
    CREATE TABLE IF NOT EXISTS events(
      id         UUID       PRIMARY KEY,
      venue      UUID       REFERENCES venues,
      start_time TIMESTAMP  WITH TIME ZONE,
      end_time   TIMESTAMP  WITH TIME ZONE,
      facebook   VARCHAR(256)                  -- ID for Facebook API
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

func main() {
  db, err := sql.Open("postgres", "host=localhost dbname=schemed sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  setup(db)
}
