package income

import (
	"context"
	"errors"
	"time"
)

type IncomeUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewIncomeUsecase(repository Repository, timeout time.Duration) Usecase {
	return &IncomeUsecase{
		Repo:           repository,
		contextTimeout: timeout,
	}
}

func (usecase *IncomeUsecase) Create(context context.Context, income IncomeDomain) (IncomeDomain, error) {
	if income.Description == "" {
		return IncomeDomain{}, errors.New("Description can't be empty")
	}
	if income.Total == 0 {
		return IncomeDomain{}, errors.New("Income can't be 0")
	}

	income, err := usecase.Repo.Create(context, income)
	if err != nil {
		return IncomeDomain{}, err
	}

	return income, nil
}
