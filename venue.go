package main

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Venue represents a place where events can occur.
type Venue struct {
	Model
	GoogleID     string `json:"googleId" form:"googleId"`
	FoursquareID string `json:"foursquareId" form:"foursquareId"`
	YelpID       string `json:"yelpId" form:"yelpId"`
}

// VenueService exposes Venue's endpoints
type VenueService struct {
	DB        *mgo.Database
	ModelName string
}

// Mount registers all Venue endpoints to the specified router.
func (s VenueService) Mount(r gin.IRouter) {
	path := "/venues"
	r.GET(path+"/:id", s.Get)
	r.GET(path, s.Fetch)
	r.POST(path, s.Create)
	r.POST(path+"/:id", s.Update)
	r.DELETE(path+"/:id", s.Delete)
}

// Get a single Venue
func (s VenueService) Get(c *gin.Context) {
	id := c.Param("id")
	var model Venue
	err := s.DB.C(s.ModelName).FindId(id).One(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": s.ModelName + " not found"})
	}
	c.JSON(http.StatusOK, model)
}

// Fetch Venues
func (s VenueService) Fetch(c *gin.Context) {
	var models []Venue
	s.DB.C(s.ModelName).Find(nil).All(&models)
	c.JSON(http.StatusOK, models)
}

// Create a Venue
func (s VenueService) Create(c *gin.Context) {
	var body Venue
	if c.Bind(&body) == nil {
		body.ID = bson.NewObjectId()
		now := time.Now()
		body.CreatedOn = now
		body.UpdatedOn = now
		err := s.DB.C(s.ModelName).Insert(&body)
		if err != nil {
			log.WithError(err).Fatal("Failed to create " + s.ModelName)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create " + s.ModelName})
		} else {
			c.JSON(http.StatusOK, body)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid " + s.ModelName})
	}
}

// Update a Venue
func (s VenueService) Update(c *gin.Context) {
	var body Venue
	if c.Bind(&body) == nil {
		id := c.Param("id")
		s.DB.C(s.ModelName).UpdateId(id, &body)
		c.JSON(http.StatusOK, body)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid " + s.ModelName})
	}
}

// Delete a Venue
func (s VenueService) Delete(c *gin.Context) {
	id := c.Param("id")
	var model Venue
	s.DB.C(s.ModelName).FindId(id).One(&model)
	s.DB.C(s.ModelName).RemoveId(id)
	c.JSON(http.StatusOK, model)
}
