package income

import (
	"net/http"
	"strconv"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	"github.com/Rahmanwghazi/Monefy/app/presenter"
	"github.com/Rahmanwghazi/Monefy/app/presenter/income/requests"
	"github.com/Rahmanwghazi/Monefy/app/presenter/income/responses"
	"github.com/Rahmanwghazi/Monefy/business/income"
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

func (incomeController IncomeController) Create(echoContext echo.Context) error {
	createIncome := requests.Income{}
	err := echoContext.Bind(&createIncome)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	income := createIncome.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	income.UserID = claims.ID

	result, err := incomeController.IncomeUseCase.CreateIncome(income)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}

func (incomeController IncomeController) GetIncome(echoContext echo.Context) error {
	createIncome := requests.Income{}
	err := echoContext.Bind(&createIncome)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	income := createIncome.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	income.UserID = claims.ID

	result, err := incomeController.IncomeUseCase.GetIncome(income)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromArrayDomain(result))
}

func (incomeController IncomeController) EditIncome(echoContext echo.Context) error {
	income := requests.Income{}
	err := echoContext.Bind(&income)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	editedIncome := income.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	editedIncome.UserID = claims.ID

	idstr := echoContext.Param("id")
	id, err := strconv.Atoi(idstr)

	result, err := incomeController.IncomeUseCase.EditIncome(editedIncome, uint(id))
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}
