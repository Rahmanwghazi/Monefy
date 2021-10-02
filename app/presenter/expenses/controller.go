package expenses

import (
	"net/http"
	"strconv"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	"github.com/Rahmanwghazi/Monefy/app/presenter"
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
	createExpense := requests.Expense{}
	err := echoContext.Bind(&createExpense)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	expense := createExpense.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	expense.UserID = claims.ID

	result, err := expenseController.ExpenseUseCase.CreateExpense(expense)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}

func (expenseController ExpenseController) GetExpenses(echoContext echo.Context) error {
	createExpense := requests.Expense{}
	err := echoContext.Bind(&createExpense)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	expense := createExpense.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	expense.UserID = claims.ID

	result, err := expenseController.ExpenseUseCase.GetExpense(expense)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromArrayDomain(result))
}

func (expenseController ExpenseController) EditExpense(echoContext echo.Context) error {
	expense := requests.Expense{}
	err := echoContext.Bind(&expense)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	editedExpense := expense.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	editedExpense.UserID = claims.ID

	idstr := echoContext.Param("id")
	id, err := strconv.Atoi(idstr)

	result, err := expenseController.ExpenseUseCase.EditExpense(editedExpense, uint(id))
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}
