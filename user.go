package main

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User contains all user-related data.
type User struct {
	Model
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// UserFilters defines possible filters on User
// TODO (Sam): Investigate filtering
type UserFilters struct {
	Email string `json:"email,omitempty" form:"email,omitempty"`
}

// UserService exposes the User model's endpoints.
type UserService struct {
	DB        *mgo.Database
	ModelName string
}

// Mount registers all User endpoints to the specified router.
func (s UserService) Mount(r gin.IRouter) {
	path := "/users"
	r.GET(path+"/:id", s.Get)
	r.GET(path, s.Fetch)
	r.POST(path, s.Create)
	r.POST(path+"/:id", s.Update)
	r.DELETE(path+"/:id", s.Delete)
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
	if c.Bind(&body) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Not a valid user",
		})
	}

	collection := s.DB.C(s.ModelName)

	// Check for existing users with the same email
	var existingUsers []User
	err := collection.Find(bson.M{"email": body.Email}).All(existingUsers)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Unable to check existing users",
		})
	}

	if len(existingUsers) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Email is taken",
		})
	}

	// Encrypt the password
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Unable to encrypt password",
		})
	}

	// Nice defaults
	now := time.Now()
	body.ID = bson.NewObjectId()
	body.Password = string(hash[:])
	body.CreatedOn = now
	body.UpdatedOn = now

	// Create the damn thing already!
	err = s.DB.C(s.ModelName).Insert(&body)
	if err != nil {
		message := "Failed to create " + s.ModelName
		log.WithError(err).Fatal(message)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": message,
		})
	}

	// Echo the created model, without the password of course
	body.Password = ""
	c.JSON(http.StatusOK, body)
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
