package users

import (
	"net/http"

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
	userSignup := requests.UserSignup{}
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
	userSignin := requests.UserSignin{}
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
