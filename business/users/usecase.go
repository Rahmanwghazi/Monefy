package users

import (
	"context"
	"errors"
	"time"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
)

type UserUsecase struct {
	JWT            middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repository Repository, timeout time.Duration, JWT middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		JWT:            JWT,
		Repo:           repository,
		contextTimeout: timeout,
	}
}

func (usecase *UserUsecase) Signup(context context.Context, user UserDomain) (UserDomain, error) {
	if user.Username == "" {
		return UserDomain{}, errors.New("Username can't be empty")
	}

	if user.Email == "" {
		return UserDomain{}, errors.New("Email can't be empty")
	}

	if user.Password == "" {
		return UserDomain{}, errors.New("Password can't be empty")
	}

	if user.FullName == "" {
		return UserDomain{}, errors.New("Fullname can't be empty")
	}

	user, err := usecase.Repo.Signup(context, user)

	if err != nil {
		return UserDomain{}, err
	}

	return user, nil
}

func (usecase *UserUsecase) Signin(context context.Context, username string, password string) (UserDomain, error) {
	if username == "" {
		return UserDomain{}, errors.New("Username can't be empty")
	}
	if password == "" {
		return UserDomain{}, errors.New("Password can't be empty")
	}

	user, err := usecase.Repo.Signin(context, username, password)
	user.Token, err = usecase.JWT.GenerateTokenJWT(int(user.ID))
	if err != nil {
		return UserDomain{}, err
	}

	return user, nil
}
