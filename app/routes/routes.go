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

	users := echoContext.Group("users", middleware.JWTWithConfig(controllerList.JWTMiddleware))
	users.PUT("/edit", controllerList.UserController.Edit)

	income := echoContext.Group("income", middleware.JWTWithConfig(controllerList.JWTMiddleware))
	income.POST("/add", controllerList.IncomeController.CreateIncome)
	income.GET("", controllerList.IncomeController.GetIncome)
	income.GET("/:id", controllerList.IncomeController.GetIncomeById)
	income.PUT("/edit/:id", controllerList.IncomeController.EditIncome)
	income.DELETE("/delete/:id", controllerList.IncomeController.DeleteIncome)

	expenses := echoContext.Group("expenses", middleware.JWTWithConfig(controllerList.JWTMiddleware))
	expenses.POST("/add", controllerList.ExpenseController.CreateExpense)
	expenses.GET("", controllerList.ExpenseController.GetExpenses)
	expenses.GET("/:id", controllerList.ExpenseController.GetExpenseById)
	expenses.PUT("/edit/:id", controllerList.ExpenseController.EditExpense)
	expenses.DELETE("/delete/:id", controllerList.ExpenseController.DeleteExpense)

	plans := echoContext.Group("plans", middleware.JWTWithConfig(controllerList.JWTMiddleware))
	plans.POST("/add", controllerList.InvestPlanController.CreatePlan)
	plans.GET("", controllerList.InvestPlanController.GetPlans)
	plans.GET("/unfinished", controllerList.InvestPlanController.GetUnfinishedPlans)
	plans.GET("/finished", controllerList.InvestPlanController.GetfinishedPlans)
	plans.PUT("/edit/:id", controllerList.InvestPlanController.EditPlan)
	plans.GET("/:id", controllerList.InvestPlanController.GetPlanById)
	plans.DELETE("/delete/:id", controllerList.InvestPlanController.DeletePlan)
}
