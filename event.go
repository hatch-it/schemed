package main

import (
	"time"
	"github.com/jinzhu/gorm"
)

// Event defines a social gathering of any sort.
type Event struct {
	gorm.Model

	Venue     string    `json:"venue" form:"venue"`
	StartTime time.Time `json:"startTime" form:"startTime"`
	EndTime   time.Time `json:"endTime" form:"endTime"`

	facebook string
}
