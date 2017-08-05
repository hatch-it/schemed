package main

import (
	"time"
)

// Event defines a social gathering of any sort.
type Event struct {
	Model
	Venue     string    `json:"venue" form:"venue"`
	StartTime time.Time `json:"startTime" form:"startTime"`
	EndTime   time.Time `json:"endTime" form:"endTime"`
	facebook  string
}
