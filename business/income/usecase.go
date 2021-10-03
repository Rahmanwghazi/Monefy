package income

import "github.com/Rahmanwghazi/Monefy/business"

type IncomeUsecase struct {
	Repo Repository
}

func NewIncomeUsecase(repository Repository) Usecase {
	return &IncomeUsecase{
		Repo: repository,
	}
}

func (usecase *IncomeUsecase) CreateIncome(income IncomeDomain) (IncomeDomain, error) {
	result, err := usecase.Repo.CreateIncome(income)
	if err != nil {
		return IncomeDomain{}, err
	}
	return result, nil
}

func (usecase *IncomeUsecase) GetIncome(income IncomeDomain) ([]IncomeDomain, error) {
	result, err := usecase.Repo.GetIncome(income)
	if err != nil {
		return []IncomeDomain{}, err
	}
	return result, nil
}

func (usecase *IncomeUsecase) GetIncomeById(income IncomeDomain, id uint) (IncomeDomain, error) {
	result, err := usecase.Repo.GetIncomeById(income, id)
	if err != nil {
		return IncomeDomain{}, err
	}
	return result, nil
}

func (usecase *IncomeUsecase) EditIncome(income IncomeDomain, id uint) (IncomeDomain, error) {
	result, err := usecase.Repo.EditIncome(income, id)
	if err != nil {
		return IncomeDomain{}, err
	}
	return result, nil
}

func (usecase *IncomeUsecase) DeleteIncome(income IncomeDomain, id uint) (string, error) {
	result, err := usecase.Repo.DeleteIncome(income, id)
	if err != nil {
		return business.ErrorInternal.Error(), err
	}
	return result, nil
}
