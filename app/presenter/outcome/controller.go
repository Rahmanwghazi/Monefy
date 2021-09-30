package outcome

import (
	"net/http"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	controllers "github.com/Rahmanwghazi/Monefy/app/presenter"
	"github.com/Rahmanwghazi/Monefy/app/presenter/outcome/requests"
	"github.com/Rahmanwghazi/Monefy/app/presenter/outcome/responses"
	"github.com/Rahmanwghazi/Monefy/business/outcome"
	"github.com/labstack/echo/v4"
)

type OutcomeController struct {
	OutcomeUseCase outcome.Usecase
}

func NewOutcomeController(outcomeUseCase outcome.Usecase) *OutcomeController {
	return &OutcomeController{
		OutcomeUseCase: outcomeUseCase,
	}
}

func (outcomeController *OutcomeController) Create(echoContext echo.Context) error {
	createOutcome := requests.CreateOutcome{}
	err := echoContext.Bind(&createOutcome)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	outcome := createOutcome.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	outcome.UserID = claims.ID

	result, err := outcomeController.OutcomeUseCase.Create(&outcome)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, http.StatusOK, responses.FromDomain(result))
}
