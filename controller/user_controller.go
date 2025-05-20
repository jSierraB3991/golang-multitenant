package controller

import (
	"net/http"

	"github.com/jSierraB3991/golang-multitenant/request"
	serviceinterface "github.com/jSierraB3991/golang-multitenant/service_interface"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService serviceinterface.UserServiceInterface
}

func NewUserController(userService serviceinterface.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (ctrl *UserController) GetAllUsers(c echo.Context) error {
	users, err := ctrl.userService.GetAllUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) SaveUser(c echo.Context) error {
	var req request.UserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if err := ctrl.userService.SaveUser(c.Request().Context(), req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}
