package server

import (
	"log"
	"os"

	"github.com/jSierraB3991/golang-multitenant/repository"
	"github.com/jSierraB3991/golang-multitenant/router"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echoServer *echo.Echo
}

func NewServer() *Server {
	return &Server{
		echoServer: echo.New(),
	}
}

func (s *Server) Start() {

	db, err := repository.NewDatabase(os.Getenv("POSTGRE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository(db, os.Getenv("SCHEMAS"))
	if err := repository.Migrations(); err != nil {
		log.Fatal(err)
	}
	s.echoServer.Use(TenantMiddleware)

	router.Routing(s.echoServer, repository)

	s.echoServer.Logger.Fatal(s.echoServer.Start(":8080"))
}
