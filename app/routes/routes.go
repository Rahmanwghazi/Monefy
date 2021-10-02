package routes

import (
	"github.com/Rahmanwghazi/Monefy/app/presenter/expenses"
	"github.com/Rahmanwghazi/Monefy/app/presenter/income"
	"github.com/Rahmanwghazi/Monefy/app/presenter/investplans"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware        middleware.JWTConfig
	UserController       users.UserController
	IncomeController     income.IncomeController
	ExpenseController    expenses.ExpenseController
	InvestPlanController investplans.InvestPlanController
}

func (controllerList *ControllerList) Routes(echoContext *echo.Echo) {
	echoContext.POST("/signup", controllerList.UserController.Signup)
	echoContext.POST("/signin", controllerList.UserController.Signin)

	withJWT := echoContext.Group("users", middleware.JWTWithConfig(controllerList.JWTMiddleware))
	withJWT.PUT("/edit", controllerList.UserController.Edit)
	withJWT.POST("/income", controllerList.IncomeController.Create)
	withJWT.GET("/income", controllerList.IncomeController.GetIncome)
	withJWT.POST("/expense", controllerList.ExpenseController.Create)
	withJWT.GET("/expenses", controllerList.ExpenseController.GetExpenses)
	withJWT.POST("/plan", controllerList.InvestPlanController.Create)
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
