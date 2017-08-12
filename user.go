package main

import (
	"time"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	log "github.com/Sirupsen/logrus"
)

// User contains all user-related data.
type User struct {
	Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// UserFilters defines possible filters on User
// TODO (Sam): Investigate filtering
type UserFilters struct {
	Email			string	`json:"email,omitempty" form:"email,omitempty"`
}

// UserService exposes the User model's endpoints
type UserService struct {
	DB   		*mgo.Database
	ModelName	string
}

// Path of the service endpoint
func (s UserService) Path() string {
	return "/users"
}

// Get a single user
func (s UserService) Get(c *gin.Context) {
	id := c.Param("id")
	var model User
	err := s.DB.C(s.ModelName).FindId(id).One(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": s.ModelName + " not found"})
	}
	model.Password = ""
	c.JSON(http.StatusOK, model)
}

// Fetch users
func (s UserService) Fetch(c *gin.Context) {
	var models []User
	s.DB.C(s.ModelName).Find(nil).All(&models)
	c.JSON(http.StatusOK, models)
}

// Create a user 
func (s UserService) Create(c *gin.Context) {
	var body User
	if c.Bind(&body) == nil {
		if body.Email == "" || body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password required"})
		} else {
			hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
			if err != nil {
				panic("Unable to encrypt password")
			}
			body.ID = bson.NewObjectId()
			body.Password = string(hash[:])
			now := time.Now()
			body.CreatedOn = now
			body.UpdatedOn = now
			err = s.DB.C(s.ModelName).Insert(&body)
			if err != nil {
				log.WithError(err).Fatal("Failed to create " + s.ModelName)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create " + s.ModelName})
			} else {
				body.Password = ""
				c.JSON(http.StatusOK, body)
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid user"})
	}
}

// Update a user
func (s UserService) Update(c *gin.Context) {
	var body User
	if c.Bind(&body) == nil {
		if body.Password != "" {
			hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
			if err != nil {
				panic("Unable to encrypt password")
			}
			body.Password = string(hash[:])
		}
		id := c.Param("id")
		s.DB.C(s.ModelName).UpdateId(id, &body)
		body.Password = ""
		c.JSON(http.StatusOK, body)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid " + s.ModelName})
	}
}

// Delete a user
func (s UserService) Delete(c *gin.Context) {
	id := c.Param("id")
	var model User
	s.DB.C(s.ModelName).FindId(id).One(&model)
	s.DB.C(s.ModelName).RemoveId(id)
	c.JSON(http.StatusOK, model)
}
