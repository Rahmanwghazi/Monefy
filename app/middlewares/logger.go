package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(echo *echo.Echo) {
	echo.Use(middleware.Logger())
}
