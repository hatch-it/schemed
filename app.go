package schemed

import (
	"github.com/gin-gonic/gin"
	"github.com/hatch-it/schemed/services"
	"gopkg.in/mgo.v2"
)

// App contains all the state for the entire application.
type App struct {
	Router   *gin.Engine
	Session  *mgo.Session
	DB       *mgo.Database
	Services []Service
}

// New creates an instance of App with the details required to connect to the database.
// Creates a connection and wire up the routes to response accordingly.
func New(hostname, dbname string) App {
	a := App{}

	var err error
	a.Session, err = mgo.Dial(hostname)
	if err != nil {
		panic(err)
	}

	a.DB = a.Session.DB(dbname)

	a.Router = gin.Default()
	a.Services = []Service{
		services.UserService{DB: a.DB},
		services.EventService{DB: a.DB},
		services.VenueService{DB: a.DB},
	}

	for _, service := range a.Services {
		service.Mount(a.Router)
	}

	return a
}

// Close all active connections running on the application.
func (a *App) Close() {
	a.Session.Close()
}

// Run starts the application.
func (a *App) Run(addr string) {
	a.Router.Run(addr)
}
