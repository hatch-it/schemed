package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

// Service is a publicly exposed endpoint for resources.
type Service interface {
	// Get an instance of a resource.
	Get(db *sql.DB, c *gin.Context) error

	// Get several instances of a resource.
	Fetch(db *sql.DB, c *gin.Context) error

	// Create an instance of a resource.
	Create(db *sql.DB, c *gin.Context) error

	// Update an existing instance of a resource.
	Update(db *sql.DB, c *gin.Context) error

	// Delete an existing instance of a resource.
	Delete(db *sql.DB, c *gin.Context) error
}

type Services []Service

var services = Services{
	User,
}
