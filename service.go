package main

import (
	"github.com/gin-gonic/gin"
)

// Service is a publicly exposed endpoint for resources.
type Service interface {
	// Get the name of the endpoint
	GetName() string

	// Initialize the service with access to the database.
	Initialize() error

	// Get an instance of a resource.
	Get(c *gin.Context)

	// Get several instances of a resource.
	Fetch(c *gin.Context)

	// Create an instance of a resource.
	Create(c *gin.Context)

	// Update an existing instance of a resource.
	Update(c *gin.Context)

	// Delete an existing instance of a resource.
	Delete(c *gin.Context)
}