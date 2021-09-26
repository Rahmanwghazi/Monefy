package routes

import (
	"github.com/Rahmanwghazi/Monefy/controllers/users"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (controllerList *ControllerList) Routes(echoContext *echo.Echo) {
	echoContext.POST("/signup", controllerList.UserController.Signup)
	echoContext.POST("/signin", controllerList.UserController.Signin)
}
