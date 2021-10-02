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

func (investPlanController InvestPlanController) Create(echoContext echo.Context) error {
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

func (investPlanController InvestPlanController) GetPlans(echoContext echo.Context) error {
	investPlanData := requests.InvestPlan{}
	err := echoContext.Bind(&investPlanData)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	investplans := investPlanData.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	investplans.UserID = claims.ID

	result, err := investPlanController.InvestPlanUseCase.GetPlans(investplans)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromArrayDomain(result))
}

func (investPlanController InvestPlanController) EditPlan(echoContext echo.Context) error {
	investPlan := requests.InvestPlan{}
	err := echoContext.Bind(&investPlan)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	editedPlan := investPlan.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	editedPlan.UserID = claims.ID

	idstr := echoContext.Param("id")
	id, err := strconv.Atoi(idstr)

	result, err := investPlanController.InvestPlanUseCase.EditPlan(editedPlan, uint(id))
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}
