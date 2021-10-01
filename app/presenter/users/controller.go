package users

import (
	"net/http"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	controllers "github.com/Rahmanwghazi/Monefy/app/presenter"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users/requests"
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

func (userController UserController) Signup(echoContext echo.Context) error {
	userSignup := requests.User{}
	err := echoContext.Bind(&userSignup)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	user, err := userController.UserUseCase.Signup(userSignup.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(echoContext, http.StatusCreated, signup.FromDomain(user))
}

func (userController UserController) Signin(echoContext echo.Context) error {
	userSignin := requests.User{}
	err := echoContext.Bind(&userSignin)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	user, err := userController.UserUseCase.Signin(userSignin.ToDomain().Username, userSignin.ToDomain().Password)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(echoContext, http.StatusOK, signin.FromDomain(user))
}

func (userController UserController) Edit(echoContext echo.Context) error {
	userEdit := requests.User{}
	err := echoContext.Bind(&userEdit)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	userEdited := userEdit.ToDomain()
	claims, err := middlewares.ExtractClaims(echoContext)
	userEdited.ID = claims.ID

	result, err := userController.UserUseCase.Edit(userEdited)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(echoContext, http.StatusCreated, signup.FromDomain(result))
}
