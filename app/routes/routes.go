package routes

import (
	"github.com/Rahmanwghazi/Monefy/controllers/income"
	"github.com/Rahmanwghazi/Monefy/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController users.UserController
	JWTMiddleware  middleware.JWTConfig

	IncomeController income.IncomeController
}

func (controllerList *ControllerList) Routes(echoContext *echo.Echo) {
	echoContext.POST("/signup", controllerList.UserController.Signup)
	echoContext.POST("/signin", controllerList.UserController.Signin)
}
