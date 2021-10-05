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

	withJWT.POST("/income", controllerList.IncomeController.CreateIncome)
	withJWT.GET("/income", controllerList.IncomeController.GetIncome)
	withJWT.GET("/income/:id", controllerList.IncomeController.GetIncomeById)
	withJWT.PUT("/income/:id", controllerList.IncomeController.EditIncome)
	withJWT.DELETE("/income/:id", controllerList.IncomeController.DeleteIncome)

	withJWT.POST("/expense", controllerList.ExpenseController.CreateExpense)
	withJWT.GET("/expenses", controllerList.ExpenseController.GetExpenses)
	withJWT.GET("/expense/:id", controllerList.ExpenseController.GetExpenseById)
	withJWT.PUT("/expense/:id", controllerList.ExpenseController.EditExpense)
	withJWT.DELETE("/expense/:id", controllerList.ExpenseController.DeleteExpense)

	withJWT.POST("/plan", controllerList.InvestPlanController.CreatePlan)
	withJWT.GET("/plans", controllerList.InvestPlanController.GetPlans)
	withJWT.GET("/plans/unfinished", controllerList.InvestPlanController.GetUnfinishedPlans)
	withJWT.GET("/plans/finished", controllerList.InvestPlanController.GetfinishedPlans)
	withJWT.PUT("/plan/:id", controllerList.InvestPlanController.EditPlan)
	withJWT.GET("/plan/:id", controllerList.InvestPlanController.GetPlanById)
	withJWT.DELETE("/plan/:id", controllerList.InvestPlanController.DeletePlan)
}
