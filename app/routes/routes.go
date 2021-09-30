package routes

import (
	"github.com/Rahmanwghazi/Monefy/app/presenter/income"
	"github.com/Rahmanwghazi/Monefy/app/presenter/outcome"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	UserController    users.UserController
	IncomeController  income.IncomeController
	OutcomeController outcome.OutcomeController
}

func (controllerList *ControllerList) Routes(echoContext *echo.Echo) {
	echoContext.POST("/signup", controllerList.UserController.Signup)
	echoContext.POST("/signin", controllerList.UserController.Signin)

	withJWT := echoContext.Group("users", middleware.JWTWithConfig(controllerList.JWTMiddleware))
	withJWT.POST("/income", controllerList.IncomeController.Create)
	withJWT.POST("/outcome", controllerList.OutcomeController.Create)
}

//edit & delete user validation
/* func UserValidation(incomeController income.IncomeController) echo.MiddlewareFunc {
	return func(handlerfunc echo.HandlerFunc) echo.HandlerFunc {
		return func(echoContext echo.Context) error {
			claims := middlewares.GetUser(echoContext)
			userId := claims.ID
			param := echoContext.Param("id")
			incomeID, _ := strconv.Atoi(param)

			userID := int(incomeController.GetById(incomeID).UserID)

			if userId == userID {
				return handlerfunc(echoContext)
			} else {
				return controllers.NewErrorResponse(echoContext, http.StatusForbidden, nil)
			}
		}
	}
} */
