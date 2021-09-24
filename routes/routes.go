package routes

import (
	"github.com/Rahmanwghazi/Monefy/constants"
	"github.com/Rahmanwghazi/Monefy/controllers"
	middleware "github.com/Rahmanwghazi/Monefy/middleware"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	echo := echo.New()
	echo.POST("/signup", controllers.SignUpController)
	echo.POST("/signin", controllers.SignInController)
	middleware.LogMiddleware(echo)

	//echoAuthBasic := echo.Group("/auth")
	//echoAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))

	echoJWT := echo.Group("/jwt")
	echoJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	//balance table

	return echo
}
