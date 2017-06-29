package models

import "time"

type Event struct {
	schemed.Model

	venue      string
	start_time time.Time
	end_time   time.Time

	facebook string
}
