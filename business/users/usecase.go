package users

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/app/middlewares"
	"github.com/Rahmanwghazi/Monefy/business"
	"github.com/Rahmanwghazi/Monefy/helpers/encrypt"
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

func (usecase *UserUsecase) Signup(user UserDomain) (UserDomain, error) {

	var errorHash error
	user.HashPassword, errorHash = encrypt.HashPassword(&user.Password)
	if errorHash != nil {
		return UserDomain{}, business.ErrorInternal
	}

	user, err := usecase.Repo.Signup(user)

	if err != nil {
		return UserDomain{}, err
	}
	return user, nil
}

func (usecase *UserUsecase) Signin(username string, password string) (UserDomain, error) {
	user, err := usecase.Repo.Signin(username)
	if err != nil {
		return UserDomain{}, business.ErrorInvalidSigninInfo
	}

	if !encrypt.CheckPasswordHash(user.HashPassword, password) {
		return UserDomain{}, business.ErrorInvalidSigninInfo
	}

	user.Token, err = usecase.JWT.GenerateTokenJWT(user.ID)
	if err != nil {
		return UserDomain{}, err
	}

	return user, nil
}

func (usecase *UserUsecase) Edit(user UserDomain) (UserDomain, error) {
	var errorHash error
	user.HashPassword, errorHash = encrypt.HashPassword(&user.Password)
	if errorHash != nil {
		return UserDomain{}, business.ErrorInternal
	}

	user, err := usecase.Repo.Edit(user)

	if err != nil {
		return UserDomain{}, err
	}
	return user, nil
}
