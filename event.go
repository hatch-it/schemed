package main

import (
	"time"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
	log "github.com/Sirupsen/logrus"
)

// Event defines a social gathering of any sort.
type Event struct {
	Model
	VenueID		bson.ObjectId	`json:"venueId" form:"venueId"`
	StartTime	time.Time 		`json:"startTime" form:"startTime"`
	EndTime		time.Time 		`json:"endTime" form:"endTime"`
	FacebookID	string			`json:"facebookId" form:"facebookId"`
}

// EventFilters contains all possible filters on Event
// TODO (Sam): Investigate filtering
type EventFilters struct {
	ModelFilters
	VenueID			*bson.ObjectId	`json:"venueId,omitempty" form:"venueId,omitempty"`
	StartTime		*time.Time		`json:"startTime,omitempty" form:"startTime,omitempty"`
	EndTime			*time.Time		`json:"endTime,omitempty" form:"endTime,omitempty"`
	FacebookID		*string			`json:"facebookId,omitempty" form:"facebookId,omitempty"`
}

// EventService exposes the Event model's endpoints
type EventService struct {
	DB			*mgo.Database
	Endpoint	string
	ModelName	string
}

// Initialize the service
func (s EventService) Initialize() string {
	return s.Endpoint
}

// Get a single Event
func (s EventService) Get(c *gin.Context) {
	id := c.Param("id")
	var model Event
	err := s.DB.C(s.ModelName).FindId(id).One(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": s.ModelName + " not found"})
	}
	c.JSON(http.StatusOK, model)
}

// Fetch Events
func (s EventService) Fetch(c *gin.Context) {
	var models []Event
	s.DB.C(s.ModelName).Find(nil).All(&models)
	c.JSON(http.StatusOK, models)
}

// Create an Event 
func (s EventService) Create(c *gin.Context) {
	var body Event
	if c.Bind(&body) == nil {
		if body.VenueID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Venue required"})
		} else {
			body.ID = bson.NewObjectId()
			now := time.Now()
			body.CreatedOn = now
			body.UpdatedOn = now
			err := s.DB.C(s.ModelName).Insert(&body)
			if err != nil {
				log.WithError(err).Fatal("Failed to create " + s.ModelName)
			}
			c.JSON(http.StatusOK, body)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid " + s.ModelName})
	}
}

// Update an Event
func (s EventService) Update(c *gin.Context) {
	var body Event
	if c.Bind(&body) == nil {
		id := c.Param("id")
		s.DB.C(s.ModelName).UpdateId(id, &body)
		c.JSON(http.StatusOK, body)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid " + s.ModelName})
	}
}

// Delete an Event
func (s EventService) Delete(c *gin.Context) {
	id := c.Param("id")
	var model Event
	s.DB.C(s.ModelName).FindId(id).One(&model)
	s.DB.C(s.ModelName).RemoveId(id)
	c.JSON(http.StatusOK, model)
}
