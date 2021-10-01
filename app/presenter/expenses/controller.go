package expenses

import (
	"net/http"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	controllers "github.com/Rahmanwghazi/Monefy/app/presenter"
	"github.com/Rahmanwghazi/Monefy/app/presenter/expenses/requests"
	"github.com/Rahmanwghazi/Monefy/app/presenter/expenses/responses"
	"github.com/Rahmanwghazi/Monefy/business/expenses"
	"github.com/labstack/echo/v4"
)

type ExpenseController struct {
	ExpenseUseCase expenses.Usecase
}

func NewExpenseController(expenseUseCase expenses.Usecase) *ExpenseController {
	return &ExpenseController{
		ExpenseUseCase: expenseUseCase,
	}
}

func (expenseController ExpenseController) Create(echoContext echo.Context) error {
	createExpense := requests.CreateExpense{}
	err := echoContext.Bind(&createExpense)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	expense := createExpense.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	expense.UserID = claims.ID

	result, err := expenseController.ExpenseUseCase.Create(expense)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}

func (expenseController ExpenseController) GetExpenses(echoContext echo.Context) error {
	createExpense := requests.CreateExpense{}
	err := echoContext.Bind(&createExpense)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	expense := createExpense.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	expense.UserID = claims.ID

	result, err := expenseController.ExpenseUseCase.GetExpense(expense)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, http.StatusOK, responses.FromArrayDomain(result))
}
