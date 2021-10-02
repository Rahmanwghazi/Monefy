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
	createInvestPlan := requests.InvestPlan{}
	err := echoContext.Bind(&createInvestPlan)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	investPlan := createInvestPlan.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	investPlan.UserID = claims.ID

	id := strconv.Itoa(investPlan.ProductID)

	result, err := investPlanController.InvestPlanUseCase.Create(id, investPlan)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return presenter.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}
