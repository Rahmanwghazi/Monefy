package users

import (
	"net/http"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	"github.com/Rahmanwghazi/Monefy/app/presenter"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users/requests"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users/responses/edit"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users/responses/signin"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users/responses/signup"
	"github.com/Rahmanwghazi/Monefy/business/users"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController *UserController) Signup(echoContext echo.Context) error {
	request := requests.User{}
	err := echoContext.Bind(&request)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	user, err := userController.UserUseCase.Signup(request.ToDomain())

	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	return presenter.NewSuccessResponse(echoContext, http.StatusCreated, signup.FromDomain(user))
}

func (userController *UserController) Signin(echoContext echo.Context) error {
	request := requests.User{}
	err := echoContext.Bind(&request)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	user, err := userController.UserUseCase.Signin(request.ToDomain().Username, request.ToDomain().Password)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	return presenter.NewSuccessResponse(echoContext, http.StatusOK, signin.FromDomain(user))
}

func (userController *UserController) Edit(echoContext echo.Context) error {
	request := requests.User{}
	err := echoContext.Bind(&request)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	userEdited := request.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	userEdited.ID = claims.ID

	result, err := userController.UserUseCase.Edit(userEdited)
	if err != nil {
		return presenter.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	return presenter.NewSuccessResponse(echoContext, http.StatusCreated, edit.FromDomain(result))
}
