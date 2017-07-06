package main

import (
	"log"
	"errors"
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

// Methods required by schemed.Service
func (user *User) Get(db *sql.DB, c *gin.Context) error {
	return errors.New("Not implemented")
}
func (user *User) Fetch(db *sql.DB, c *gin.Context) error {
	return errors.New("Not implemented")
}
func (user *User) Create(db *sql.DB, c *gin.Context) error {
	return errors.New("Not implemented")
}
func (user *User) Update(db *sql.DB, c *gin.Context) error {
	return errors.New("Not implemented")
}
func (user *User) Delete(db *sql.DB, c *gin.Context) error {
	return errors.New("Not implemented")
}
