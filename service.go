package main

import (
	"github.com/gin-gonic/gin"
)

// Service is a publicly exposed endpoint for resources.
type Service interface {
	// Mount the service (and all of its endpoints) to a router
	Mount(r gin.IRouter)
}
