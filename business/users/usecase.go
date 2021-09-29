package users

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
)

type UserUsecase struct {
	JWT            *middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repository Repository, JWT *middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		JWT:  JWT,
		Repo: repository,
	}
}

func (usecase *UserUsecase) Signup(user *UserDomain) (UserDomain, error) {
	result, err := usecase.Repo.Signup(user)
	if err != nil {
		return UserDomain{}, err
	}
	return result, nil
}

func (usecase *UserUsecase) Signin(username string, password string) (UserDomain, error) {
	user, err := usecase.Repo.Signin(username, password)
	user.Token, err = usecase.JWT.GenerateTokenJWT(user.ID)
	if err != nil {
		return UserDomain{}, err
	}

	return user, nil
}
