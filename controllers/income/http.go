package income

import (
	"net/http"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
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

func (incomeController *IncomeController) Create(echoContext echo.Context) error {
	createIncome := requests.CreateIncome{}
	err := echoContext.Bind(&createIncome)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	userId := uint(middlewares.GetUser(echoContext).ID)
	income := createIncome.ToDomain()
	result, err := incomeController.IncomeUseCase.Create(userId, income)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}
