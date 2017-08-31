package services

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hatch-it/schemed/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserService exposes the User model's endpoints.
type UserService struct {
	DB *mgo.Database
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
	var user models.User
	err := s.DB.C("User").FindId(id).One(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User" + " not found"})
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

// Fetch users
func (s UserService) Fetch(c *gin.Context) {
	var users []models.User
	s.DB.C("User").Find(nil).All(&users)
	c.JSON(http.StatusOK, users)
}

// Create a user
func (s UserService) Create(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a valid user",
			"error":   err.Error(),
		})
		return
	}

	// Check for existing users with the same email
	count, err := s.DB.C("User").Find(bson.M{"email": user.Email}).Count()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to check existing users",
			"error":   err.Error(),
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a valid user",
			"error":   "Email has been taken already",
		})
		return
	}

	// Encrypt the password
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to encrypt password",
			"error":   err.Error(),
		})
		return
	}

	// Nice defaults
	now := time.Now()
	user.ID = bson.NewObjectId()
	user.Password = string(hash[:])
	user.CreatedOn = now
	user.UpdatedOn = now

	// Create the damn thing already!
	if err := s.DB.C("User").Insert(&user); err != nil {
		log.WithError(err).Fatal("Failed to create User")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create User",
			"error":   err.Error(),
		})
		return
	}

	// Echo the created model, without the password of course
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

// Update a user
func (s UserService) Update(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a valid User",
			"error":   err.Error(),
		})
		return
	}

	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to encrypt password",
				"error":   err.Error(),
			})
			return
		}
		user.Password = string(hash[:])
	}
	id := c.Param("id")
	s.DB.C("User").UpdateId(id, &user)
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

// Delete a user
func (s UserService) Delete(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	s.DB.C("User").FindId(id).One(&user)
	s.DB.C("User").RemoveId(id)
	c.JSON(http.StatusOK, user)
}
