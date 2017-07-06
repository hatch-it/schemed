package main

import "time"

// Event defines a social gathering of any sort.
type Event struct {
	Model

	Venue     string    `json:"venue"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`

	facebook string
}
