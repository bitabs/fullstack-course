package server

import (
	"fullstack-course/api"
	. "fullstack-course/db"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	DB 				*DB
	Router			*gin.Engine
	Environment		string
	PORT			string
}

// serve our react frontend
func (s Server) serveFrontend() {
	if s.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
		s.Router.Use(static.Serve("/", static.LocalFile("./build", true)))
	}
}

func (s Server) serve() {
	// log which port we're connected to
	log.Println("Now server is running on port" + s.PORT)

	// initiate the app on available port
	if err := s.Router.Run(s.PORT); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Server() {

	s.initAPI()

	s.serveFrontend()

	s.serve()

}

func (s *Server) initAPI() {
	api.InitAPI(s.DB, s.Router, s.Environment)
}
