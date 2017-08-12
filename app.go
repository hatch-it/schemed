
package main

import (
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
)

// App contains all the state for the entire application.
type App struct {
	Router 		*gin.Engine
	Session		*mgo.Session
	DB 			*mgo.Database
	Services 	[]Service
}

// Initialize takes the details required to connect to the database.
// Create a connection and wire up the routes to response accordingly.
func (a *App) Initialize(hostname, dbname string) {
	var err error
	a.Session, err = mgo.Dial(hostname)
	if err != nil {
		panic(err)
	}

	a.DB = a.Session.DB(dbname)

	a.Router = gin.Default()
	a.Services = []Service{
		UserService{a.DB, "User"},
		EventService{a.DB, "Event"},
		VenueService{a.DB, "Venue"},
	}

	for _, service := range a.Services {
		path := service.Path()
		a.Router.GET(path + "/:id", service.Get)
		a.Router.GET(path, service.Fetch)
		a.Router.POST(path, service.Create)
		a.Router.POST(path + "/:id", service.Update)
		a.Router.DELETE(path + "/:id", service.Delete)
	}
}

// Close all active connections running on the application.
func (a *App) Close() {
	a.Session.Close()
}

// Run starts the application.
func (a *App) Run(addr string) {
	a.Router.Run(addr)
}
