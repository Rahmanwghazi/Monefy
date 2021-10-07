package investplans

import (
	"net/http"
	"strconv"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	"github.com/Rahmanwghazi/Monefy/app/presenter"
	"github.com/Rahmanwghazi/Monefy/app/presenter/investplans/requests"
	"github.com/Rahmanwghazi/Monefy/app/presenter/investplans/responses"
	"github.com/Rahmanwghazi/Monefy/business/investplans"
	"github.com/labstack/echo/v4"
)

type InvestPlanController struct {
	InvestPlanUseCase investplans.Usecase
}

func NewInvestPlanController(investPlanUseCase investplans.Usecase) *InvestPlanController {
	return &InvestPlanController{
		InvestPlanUseCase: investPlanUseCase,
	}
}

func (investPlanController *InvestPlanController) CreatePlan(echoContext echo.Context) error {
	request := requests.InvestPlan{}

	if err := echoContext.Bind(&request); err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	claims, err := middlewares.ExtractClaims(echoContext)
	request.UserID = claims.ID

	idProduct := strconv.Itoa(request.ProductID)

	response, err := investPlanController.InvestPlanUseCase.Create(idProduct, request.ToDomain())
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(response))
}

func (investPlanController *InvestPlanController) GetPlans(echoContext echo.Context) error {
	request := requests.InvestPlan{}

	investplans := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	investplans.UserID = claims.ID

	result, err := investPlanController.InvestPlanUseCase.GetPlans(investplans)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromArrayDomain(result))
}

func (investPlanController *InvestPlanController) GetUnfinishedPlans(echoContext echo.Context) error {
	request := requests.InvestPlan{}

	investplans := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	investplans.UserID = claims.ID

	result, err := investPlanController.InvestPlanUseCase.GetUnfinishedPlans(investplans)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromArrayDomain(result))
}

func (investPlanController *InvestPlanController) GetfinishedPlans(echoContext echo.Context) error {
	request := requests.InvestPlan{}

	investplans := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	investplans.UserID = claims.ID

	result, err := investPlanController.InvestPlanUseCase.GetfinishedPlans(investplans)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromArrayDomain(result))
}

func (investPlanController *InvestPlanController) GetPlanById(echoContext echo.Context) error {
	request := requests.InvestPlan{}

	investplans := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	investplans.UserID = claims.ID

	idParam := echoContext.Param("id")
	id, err := strconv.Atoi(idParam)

	result, err := investPlanController.InvestPlanUseCase.GetPlanById(investplans, uint(id))
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}

func (investPlanController InvestPlanController) EditPlan(echoContext echo.Context) error {
	request := requests.InvestPlan{}
	err := echoContext.Bind(&request)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	editedPlan := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	editedPlan.UserID = claims.ID

	idParam := echoContext.Param("id")
	id, err := strconv.Atoi(idParam)

	result2, err2 := investPlanController.InvestPlanUseCase.EditPlan(editedPlan, uint(id))
	if err2 != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err2)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result2))
}

func (investPlanController *InvestPlanController) DeletePlan(echoContext echo.Context) error {
	request := requests.InvestPlan{}

	editedPlan := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	editedPlan.UserID = claims.ID

	idParam := echoContext.Param("id")
	id, err := strconv.Atoi(idParam)

	result, err := investPlanController.InvestPlanUseCase.GetPlanById(editedPlan, uint(id))

	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusNotFound, err)
	}

	result.UserID = claims.ID

	result2, err2 := investPlanController.InvestPlanUseCase.DeletePlan(&result, uint(id))
	if err2 != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err2)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, result2)
}
