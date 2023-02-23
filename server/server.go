package server

import (
	"log"
	"ps-tag-onboarding-go/db"
	"ps-tag-onboarding-go/services/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() {
	// New DB based on config. For now we use mongo.
	mongo, err := db.NewMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	userService, err := user.NewUserService(mongo)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET(user.RouteUser, userService.HandleUserGet)
	e.POST(user.RouteUser, userService.HandleUserSet)

	e.Logger.Fatal(e.Start(":8080"))
}
