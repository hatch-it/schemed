package main

import (
	"time"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User contains all user-related data.
type User struct {
	Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Users []User

// UserService exposes the User model's endpoints
type UserService struct {
	DB   *mgo.Database
	Name string
}

// Methods required by schemed.Service
func (s UserService) Initialize() string {
	return "users"
}
func (s UserService) Get(c *gin.Context) {
	id := c.Param("id")
	var model User
	err := s.DB.C("User").FindId(id).One(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
	model.Password = ""
	c.JSON(http.StatusOK, model)
}
func (s UserService) Fetch(c *gin.Context) {
	var filters User
	if c.Bind(&filters) == nil {
		filters.Password = "" // Don't filter by password
		var models Users
		s.DB.C("User").Find(&filters).All(&models)
		c.JSON(http.StatusOK, models)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Filters not recognized"})
	}
}
func (s UserService) Create(c *gin.Context) {
	var body User
	if c.BindJSON(&body) == nil {
		if body.Email == "" || body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password required"})
		} else {
			hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
			if err != nil {
				panic("Unable to encrypt password")
			}
			body.Password = string(hash[:])
			now := time.Now()
			body.CreatedOn = now
			body.UpdatedOn = now
			s.DB.C("User").Insert(&body)
			body.Password = ""
			c.JSON(http.StatusOK, body)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid user"})
	}
}
func (s UserService) Update(c *gin.Context) {
	var body User
	if c.BindJSON(&body) == nil {
		if body.Password != "" {
			hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
			if err != nil {
				panic("Unable to encrypt password")
			}
			body.Password = string(hash[:])
		}
		id := c.Param("id")
		s.DB.C("Users").UpdateId(id, &body)
		body.Password = ""
		c.JSON(http.StatusOK, body)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid user"})
	}
}
func (s UserService) Delete(c *gin.Context) {
	id := c.Param("id")
	var model User
	s.DB.C("User").FindId(id).One(&model)
	s.DB.C("User").RemoveId(id)
	c.JSON(http.StatusOK, model)
}
