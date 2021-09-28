package income

import (
	"net/http"

	"github.com/Rahmanwghazi/Monefy/business/income"
	"github.com/Rahmanwghazi/Monefy/controllers"
	"github.com/Rahmanwghazi/Monefy/controllers/income/requests"
	"github.com/Rahmanwghazi/Monefy/controllers/income/responses"
	"github.com/labstack/echo/v4"
)

type IncomeController struct {
	IncomeUseCase income.Usecase
}

func NewIncomeController(incomeUseCase income.Usecase) *IncomeController {
	return &IncomeController{
		IncomeUseCase: incomeUseCase,
	}
}

func (incomeController IncomeController) Signup(echoContext echo.Context) error {
	createIncome := requests.CreateIncome{}
	err := echoContext.Bind(&createIncome)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	context := echoContext.Request().Context()
	user, err := incomeController.IncomeUseCase.Create(context, createIncome.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(echoContext, http.StatusCreated, responses.FromDomain(user))
}
