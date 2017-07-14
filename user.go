package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User contains all user-related data.
type User struct {
	gorm.Model

	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Users []User

// UserService exposes the User model's endpoints
type UserService struct {
	DB   *gorm.DB
	Name string
}

// Methods required by schemed.Service
func (s UserService) GetName() string {
	return s.Name
}
func (s UserService) Initialize() {
	s.DB.AutoMigrate(&User{})
}
func (s UserService) Get(c *gin.Context) {
	id := c.Param("id")
	var model User
	s.DB.First(&model, id)
	model.Password = ""
	c.JSON(http.StatusOK, model)
}
func (s UserService) Fetch(c *gin.Context) {
	var filters User
	if c.Bind(&filters) == nil {
		filters.Password = "" // Don't filter by password
		var models Users
		s.DB.Where(&filters).Find(&models)
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
			s.DB.Create(&body)
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
		s.DB.Where("id = ?", id).Updates(&body)
		body.Password = ""
		c.JSON(http.StatusOK, body)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid user"})
	}
}
func (s UserService) Delete(c *gin.Context) {
	id := c.Param("id")
	var model User
	s.DB.First(&model, id)
	s.DB.Model(&model).Delete(&User{})
	c.JSON(http.StatusOK, model)
}
