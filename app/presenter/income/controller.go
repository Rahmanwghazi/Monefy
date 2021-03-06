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

func (incomeController *IncomeController) CreateIncome(echoContext echo.Context) error {
	request := requests.Income{}
	err := echoContext.Bind(&request)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	income := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	income.UserID = claims.ID

	result, err := incomeController.IncomeUseCase.CreateIncome(income)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}

func (incomeController *IncomeController) GetIncome(echoContext echo.Context) error {
	request := requests.Income{}

	income := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	income.UserID = claims.ID

	result, err := incomeController.IncomeUseCase.GetIncome(income)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromArrayDomain(result))
}

func (incomeController *IncomeController) GetIncomeById(echoContext echo.Context) error {
	request := requests.Income{}

	income := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	income.UserID = claims.ID

	idParam := echoContext.Param("id")
	id, err := strconv.Atoi(idParam)

	result, err := incomeController.IncomeUseCase.GetIncomeById(income, uint(id))
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}

func (incomeController *IncomeController) EditIncome(echoContext echo.Context) error {
	request := requests.Income{}
	err := echoContext.Bind(&request)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	editedIncome := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	editedIncome.UserID = claims.ID

	idParam := echoContext.Param("id")
	id, err := strconv.Atoi(idParam)

	result2, err2 := incomeController.IncomeUseCase.EditIncome(editedIncome, uint(id))
	if err2 != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err2)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result2))
}

func (incomeController *IncomeController) DeleteIncome(echoContext echo.Context) error {
	request := requests.Income{}

	editedIncome := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	editedIncome.UserID = claims.ID

	idParam := echoContext.Param("id")
	id, err := strconv.Atoi(idParam)

	result, err := incomeController.IncomeUseCase.GetIncomeById(editedIncome, uint(id))

	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusNotFound, err)
	}

	result2, err2 := incomeController.IncomeUseCase.DeleteIncome(&result, uint(id))
	if err2 != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err2)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, result2)
}
