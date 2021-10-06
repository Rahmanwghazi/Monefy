package users_test

import (
	"testing"
	"time"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	"github.com/Rahmanwghazi/Monefy/business/users"
	_userMock "github.com/Rahmanwghazi/Monefy/business/users/mocks"
	"github.com/Rahmanwghazi/Monefy/helpers/encrypt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockUserRepository _userMock.Repository
	userDomain         users.UserDomain
	userUsecase        users.Usecase
)

func TestMain(m *testing.M) {
	userUsecase = users.NewUserUsecase(&mockUserRepository, &middlewares.ConfigJWT{})
	hashPassword, _ := encrypt.HashPassword("123")
	userDomain = users.UserDomain{
		ID:           1,
		Username:     "wafiq",
		Email:        "wafiq@monefy.com",
		HashPassword: hashPassword,
		Fullname:     "rahman wafiq ghazi",
		Dob:          time.Now(),
		Token:        "token",
	}
	m.Run()
}

func TestSignup(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockUserRepository.On("Signup", mock.Anything).Return(userDomain, nil).Once()

		input := users.UserDomain{
			ID:       1,
			Username: "wafiq",
			Email:    "wafiq@monefy.com",
			Password: "123",
			Fullname: "rahman wafiq ghazi",
			Dob:      time.Now(),
		}

		result, err := userUsecase.Signup(&input)
		assert.Nil(t, err)
		assert.Equal(t, userDomain, result)
	})

	t.Run("Test case 2 - Invalid (duplicate email)", func(t *testing.T) {
		mockUserRepository.On("Signup", mock.Anything).Return(users.UserDomain{}, assert.AnError).Once()

		input := users.UserDomain{
			ID:       1,
			Username: "wafiq2",
			Email:    "wafiq@monefy.com",
			Password: "123",
			Fullname: "rahman wafiq ghazi",
			Dob:      time.Now(),
		}

		result, err := userUsecase.Signup(&input)
		assert.NotNil(t, err)
		assert.NotEqual(t, userDomain, result)
	})
}

func TestSignin(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockUserRepository.On("Signin", mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		input := users.UserDomain{
			Username: "wafiq",
			Password: "123",
		}

		result, err := userUsecase.Signin(input.Username, input.Password)
		result.Token = "token"
		assert.Nil(t, err)
		assert.Equal(t, userDomain, result)
	})

	t.Run("Test case 2 - Invalid (wrong username/password)", func(t *testing.T) {
		mockUserRepository.On("Signin", mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		input := users.UserDomain{
			Username: "wafiq",
			Password: "2123",
		}

		result, err := userUsecase.Signin(input.Username, input.Password)
		result.Token = "token"
		assert.NotNil(t, err)
		assert.NotEqual(t, userDomain, result)
	})
}

func TestEdit(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockUserRepository.On("Edit", mock.Anything).Return(userDomain, nil).Once()

		input := users.UserDomain{
			ID:       1,
			Username: "wafiqedited",
			Email:    "wafiqedited@monefy.com",
			Password: "123",
			Fullname: "rahman wafiq ghazi",
			Dob:      time.Now(),
		}

		result, err := userUsecase.Edit(&input)
		assert.Nil(t, err)
		assert.Equal(t, userDomain, result)
	})

	t.Run("Test case 2 - Invalid (Duplicate email)", func(t *testing.T) {
		mockUserRepository.On("Edit", mock.Anything).Return(userDomain, assert.AnError).Once()

		input := users.UserDomain{
			ID:       2,
			Username: "wafiqedited",
			Email:    "wafiq@monefy.com",
			Password: "123",
			Fullname: "rahman wafiq ghazi",
			Dob:      time.Now(),
		}

		result, err := userUsecase.Edit(&input)
		assert.NotNil(t, err)
		assert.NotEqual(t, userDomain, result)
	})
}

//coverage 84.6%
