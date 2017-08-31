package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Event defines a social gathering of any sort.
type Event struct {
	Model
	VenueID    bson.ObjectId `json:"venueId" form:"venueId" binding:"required"`
	StartTime  time.Time     `json:"startTime" form:"startTime"`
	EndTime    time.Time     `json:"endTime" form:"endTime"`
	FacebookID string        `json:"facebookId" form:"facebookId"`
}
