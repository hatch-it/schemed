package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User contains all user-related data.
type User struct {
	Model

	Email    string `json:"email" binding:"required"`
	Password string `json:"-" binding:"required"`
}

type Users []User

// UserService exposes the User model's endpoints
type UserService struct {
	DB   *sql.DB
	Name string
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
	return err
}
func (s UserService) Get(c *gin.Context) {
	id := c.Param("id")

	var model User
	row := s.DB.QueryRow("SELECT * FROM $1 WHERE id = $2", s.Name, id)
	err := row.Scan(&model)

	c.Header("Content-Type", "application/vnd.api+json")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": []gin.H{
				gin.H{
					"status": http.StatusNotFound,
					"source": gin.H{"pointer": "/data"},
					"detail": "Could not find a User with id " + id,
				},
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"links": gin.H{
				"self": c.Request.URL,
			},
			"data": gin.H{
				"type":       s.Name,
				"id":         id,
				"attributes": model,
			},
		})
	}
}
func (s UserService) Fetch(c *gin.Context) {
	rows, err := s.DB.Query("SELECT * FROM $1", s.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	defer rows.Close()

	users := make([]User, 0)
	var user User

	for rows.Next() {
		err := rows.Scan(&user)
		if err != nil {
			log.Fatal(err) // How would you even get here?
		}
		users = append(users, user)
	}

	c.Header("Content-Type", "application/vnd.api+json")

	c.JSON(http.StatusOK, gin.H{
		"links": gin.H{
			"self": c.Request.URL,
		},
		"data": users,
	})
}

func (s UserService) Create(c *gin.Context) {
	var user User
	if c.BindJSON(&user) != nil {
		c.Header("Content-Type", "application/vnd.api+json")
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": []gin.H{
				gin.H{
					"status": http.StatusNotAcceptable,
					"title":  "Required User fields not specified",
					"detail": "The required fields of User are email and password.",
				},
			},
		})
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err) // How would you even get here?
		}
		_, err = s.DB.Exec(`INSERT INTO $1 (email, password) VALUES ('$2', '$3')`, s.Name, user.Email, hash)
		if err != nil {
			log.Fatal(err) // How would you even get here?
		}
	}
}
func (s UserService) Update(c *gin.Context) {
	var user User
	id := c.Param("id")
	if c.BindJSON(&user) != nil || id == "" {

	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err) // How would you even get here?
		}
		fields := 
		_, err = s.DB.Exec(`UPDATE $1 SET $2 WHERE id = '$3'`, s.Name, fields, id)
		if err != nil {
			log.Fatal(err) // How would you even get here?
		}
	}
}
func (s UserService) Delete(c *gin.Context) {
	c.Header("Content-Type", "application/vnd.api+json")
	c.JSON(http.StatusNotImplemented, gin.H{
		"errors": []gin.H{
			gin.H{
				"title":  "Feature not implemented",
				"detail": "We have yet to add this feature in yet. Please wait while we're braining.",
			},
		},
	})
}
