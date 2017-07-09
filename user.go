package main

import (
	"log"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User contains all user-related data.
type User struct {
	Model

	Email    string `json:"email"`
	Password string `json:"password"`
}

// Hash a password
func Hash(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

// UserService exposes the User model's endpoints
type UserService struct {
	DB 		*sql.DB
	Name 	string
}

// Methods required by schemed.Service
func (s UserService) GetName() string {
	return s.Name
}
func (s UserService) Initialize() error {
	_, err := s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id         UUID         PRIMARY KEY     DEFAULT uuid_generate_v4(),
			created_at TIMESTAMP    WITH TIME ZONE  DEFAULT now(),
			updated_at TIMESTAMP    WITH TIME ZONE  DEFAULT now(),
			deleted_at TIMESTAMP	WITH TIME ZONE,

			email      VARCHAR(254) NOT NULL,
			password   VARCHAR(74)  NOT NULL
		);
	`)
	return err;
}
func (s UserService) Get(c *gin.Context) {
	id := c.Param("id")

	var model User
	row := s.DB.QueryRow("SELECT * FROM $1 WHERE id = $2", s.Name, id)
	err := row.Scan(&model)

	c.Header("Content-Type", "application/vnd.api+json")

	var data gin.H

	if err != nil {
		data = gin.H{
			"type": s.Name,
			"id": id,
			"attributes": model,
		}
	} else {
		data = gin.H{}
	}

	c.JSON(200, gin.H{
		"data": data,
		"meta": gin.H{},
	})
}
func (s UserService) Fetch(c *gin.Context) {
	c.Header("Content-Type", "application/vnd.api+json")
	c.JSON(500, gin.H{
		"errors": []gin.H{
			gin.H{
				"title": "Feature not implemented",
				"detail": "We have yet to add this feature in yet. Please wait while we're braining.",
			},
		},
		"meta": gin.H{},
	})
}
func (s UserService) Create(c *gin.Context) {
	s.DB.Exec(`
		INSERT INTO 
	`)
	c.Header("Content-Type", "application/vnd.api+json")
	c.JSON(500, gin.H{
		"errors": []gin.H{
			gin.H{
				"title": "Feature not implemented",
				"detail": "We have yet to add this feature in yet. Please wait while we're braining.",
			},
		},
		"meta": gin.H{},
	})
}
func (s UserService) Update(c *gin.Context) {
	c.Header("Content-Type", "application/vnd.api+json")
	c.JSON(500, gin.H{
		"errors": []gin.H{
			gin.H{
				"title": "Feature not implemented",
				"detail": "We have yet to add this feature in yet. Please wait while we're braining.",
			},
		},
		"meta": gin.H{},
	})
}
func (s UserService) Delete(c *gin.Context) {
	c.Header("Content-Type", "application/vnd.api+json")
	c.JSON(500, gin.H{
		"errors": []gin.H{
			gin.H{
				"title": "Feature not implemented",
				"detail": "We have yet to add this feature in yet. Please wait while we're braining.",
			},
		},
		"meta": gin.H{},
	})
}
