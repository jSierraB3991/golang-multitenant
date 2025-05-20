package router

import (
	"github.com/jSierraB3991/golang-multitenant/controller"
	"github.com/jSierraB3991/golang-multitenant/repository"
	"github.com/jSierraB3991/golang-multitenant/service"
	"github.com/labstack/echo/v4"
)

func Routing(echoServer *echo.Echo, repo *repository.Repository) {

	userService := service.NewUserService(repo)
	userController := controller.NewUserController(userService)

	UserRouter := echoServer.Group("/user")

	UserRouter.GET("/", userController.GetAllUsers)
	UserRouter.POST("/", userController.SaveUser)
}
