package services

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hatch-it/schemed/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// VenueService exposes Venue's endpoints
type VenueService struct {
	DB *mgo.Database
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
	var venue models.Venue
	err := s.DB.C("Venue").FindId(id).One(&venue)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Venue" + " not found"})
	}
	c.JSON(http.StatusOK, venue)
}

// Fetch Venues
func (s VenueService) Fetch(c *gin.Context) {
	var venues []models.Venue
	s.DB.C("Venue").Find(nil).All(&venues)
	c.JSON(http.StatusOK, venues)
}

// Create a Venue
func (s VenueService) Create(c *gin.Context) {
	var venue models.Venue
	if c.Bind(&venue) == nil {
		venue.ID = bson.NewObjectId()
		now := time.Now()
		venue.CreatedOn = now
		venue.UpdatedOn = now
		err := s.DB.C("Venue").Insert(&venue)
		if err != nil {
			log.WithError(err).Fatal("Failed to create " + "Venue")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create " + "Venue"})
		} else {
			c.JSON(http.StatusOK, venue)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid " + "Venue"})
	}
}

// Update a Venue
func (s VenueService) Update(c *gin.Context) {
	var venue models.Venue
	if c.Bind(&venue) == nil {
		id := c.Param("id")
		s.DB.C("Venue").UpdateId(id, &venue)
		c.JSON(http.StatusOK, venue)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid " + "Venue"})
	}
}

// Delete a Venue
func (s VenueService) Delete(c *gin.Context) {
	id := c.Param("id")
	var venue models.Venue
	s.DB.C("Venue").FindId(id).One(&venue)
	s.DB.C("Venue").RemoveId(id)
	c.JSON(http.StatusOK, venue)
}
