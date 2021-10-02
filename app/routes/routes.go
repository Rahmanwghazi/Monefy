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
	withJWT.PUT("/income/:id", controllerList.IncomeController.EditIncome)

	withJWT.POST("/expense", controllerList.ExpenseController.Create)
	withJWT.GET("/expenses", controllerList.ExpenseController.GetExpenses)
	withJWT.PUT("/expense/:id", controllerList.ExpenseController.EditExpense)

	withJWT.POST("/plan", controllerList.InvestPlanController.Create)
	withJWT.GET("/plans", controllerList.InvestPlanController.GetPlans)
	withJWT.PUT("/plan/:id", controllerList.InvestPlanController.EditPlan)

	//todo: add user validation for edit by id services
}
