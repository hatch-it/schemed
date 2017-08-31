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

// EventService exposes the Event model's endpoints
type EventService struct {
	DB *mgo.Database
}

// Mount registers all Event endpoints to the specified router.
func (s EventService) Mount(r gin.IRouter) {
	path := "/events"
	r.GET(path+"/:id", s.Get)
	r.GET(path, s.Fetch)
	r.POST(path, s.Create)
	r.POST(path+"/:id", s.Update)
	r.DELETE(path+"/:id", s.Delete)
}

// Get a single Event
func (s EventService) Get(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	err := s.DB.C("Event").FindId(id).One(&event)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event" + " not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

// Fetch Events
func (s EventService) Fetch(c *gin.Context) {
	var events []models.Event
	s.DB.C("Event").Find(nil).All(&events)
	c.JSON(http.StatusOK, events)
}

// Create an Event
func (s EventService) Create(c *gin.Context) {
	var event models.Event
	if err := c.Bind(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a valid Event",
			"error":   err.Error(),
		})
		return
	}

	event.ID = bson.NewObjectId()
	now := time.Now()
	event.CreatedOn = now
	event.UpdatedOn = now
	if err := s.DB.C("Event").Insert(&event); err != nil {
		log.WithError(err).Fatal("Failed to create Event")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create Event",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, event)
}

// Update an Event
func (s EventService) Update(c *gin.Context) {
	var event models.Event

	if err := c.Bind(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid " + "Event"})
		return
	}

	id := c.Param("id")
	s.DB.C("Event").UpdateId(id, &event)
	c.JSON(http.StatusOK, event)
}

// Delete an Event
func (s EventService) Delete(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	s.DB.C("Event").FindId(id).One(&event)
	s.DB.C("Event").RemoveId(id)
	c.JSON(http.StatusOK, event)
}
