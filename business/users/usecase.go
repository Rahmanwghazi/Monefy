package users

import (
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repository Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		Repo:           repository,
		contextTimeout: timeout,
	}
}

func (usecase *UserUsecase) Signup(context context.Context, user Domain) (Domain, error) {
	if user.Username == "" {
		return Domain{}, errors.New("Username can't be empty")
	}

	if user.Email == "" {
		return Domain{}, errors.New("Email can't be empty")
	}

	if user.Password == "" {
		return Domain{}, errors.New("Password can't be empty")
	}

	if user.FullName == "" {
		return Domain{}, errors.New("Fullname can't be empty")
	}

	user, err := usecase.Repo.Signup(context, user)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (usecase *UserUsecase) Signin(context context.Context, username string, password string) (Domain, error) {
	if username == "" {
		return Domain{}, errors.New("Username can't be empty")
	}
	if password == "" {
		return Domain{}, errors.New("Password can't be empty")
	}

	user, err := usecase.Repo.Signin(context, username, password)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
