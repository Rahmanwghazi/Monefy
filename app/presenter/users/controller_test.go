package users_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Rahmanwghazi/Monefy/app/presenter"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users/responses/signin"
	"github.com/Rahmanwghazi/Monefy/app/presenter/users/responses/signup"
	"github.com/Rahmanwghazi/Monefy/business"
	_user "github.com/Rahmanwghazi/Monefy/business/users"
	_userMock "github.com/Rahmanwghazi/Monefy/business/users/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockUserUsecase       _userMock.Usecase
	userController        *users.UserController
	userDomain            _user.UserDomain
	userSigninResponse    signin.UserSignin
	userSignupResponse    signup.UserSignup
	userSigninRequest     string
	userSignupRequest     string
	invaldiBindingRequest string
)

func TestMain(m *testing.M) {
	userController = users.NewUserController(&mockUserUsecase)
	userSigninRequest = `{
		"username" : "wafiq",
		"password" : "123456"
	}`
	userSignupRequest = `{
		"username" : "wafiq",
		"password" : "123456",
		"email"    : "wafiq@gmail.com",
		"fullname" : "rahman wafiq",
		"dob"      : "2000-06-01T00:00:00Z"
	}`
	invaldiBindingRequest = `{
		"username" : "wafiq"
		"password" : "123456"
		"email"    : "wafiq@gmail.com"
		"fullname" : "rahman wafiq"
		"dob"      : "2000-06-01T00:00:00Z"
	}`
	userSigninResponse = signin.UserSignin{
		ID:        1,
		Username:  "wafiq",
		Email:     "wafiq@monefy.com",
		FullName:  "rahman wafiq",
		DoB:       time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userSignupResponse = signup.UserSignup{
		ID:        1,
		Username:  "wafiq",
		Email:     "wafiq@monefy.com",
		Fullname:  "rahman wafiq",
		DoB:       time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userDomain = _user.UserDomain{
		ID:        1,
		Username:  "wafiq",
		Email:     "wafiq@monefy.com",
		Fullname:  "rahman wafiq",
		Dob:       time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m.Run()
}

func TestSignup(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(userSignupRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		mockUserUsecase.On("Signup", mock.Anything).Return(userDomain, nil).Once()

		resp := presenter.BaseResponse{}
		resp.Meta.Status = http.StatusCreated
		resp.Meta.Message = "success"
		resp.Data = userSignupResponse
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, userController.Signup(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})

	t.Run("Test case 2 - Invalid (Duplicate email)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(userSignupRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		mockUserUsecase.On("Signup", mock.Anything).Return(userDomain, business.ErrorDuplicateEmail).Once()

		resp := presenter.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Email has already been taken"
		resp.Data = nil
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, userController.Signup(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})

	t.Run("Test case 3 - Invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(invaldiBindingRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		resp := presenter.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "code=400, message=Syntax error: offset=28, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"
		resp.Data = nil
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, userController.Signup(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})

}

func TestSignin(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/signin", strings.NewReader(userSigninRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		mockUserUsecase.On("Signin", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		resp := presenter.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "success"
		resp.Data = userSigninResponse
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, userController.Signin(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})

	t.Run("Test case 2 - Invalid", func(t *testing.T) {
		request := `{"username": "wafiq123","password": "11111"}`
		req := httptest.NewRequest(http.MethodPost, "/signin", strings.NewReader(request))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		mockUserUsecase.On("Signin", mock.Anything, mock.Anything).Return(userDomain, business.ErrorInvalidSigninInfo).Once()

		resp := presenter.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Username or password is invalid"
		resp.Data = nil
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, userController.Signin(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})

	t.Run("Test case 3 - Invalid (Binding error)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/signin", strings.NewReader(invaldiBindingRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		mockUserUsecase.On("Signin", mock.Anything, mock.Anything).Return(userDomain, business.ErrorInternal).Once()

		resp := presenter.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "code=400, message=Syntax error: offset=28, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"
		resp.Data = nil
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, userController.Signin(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}
