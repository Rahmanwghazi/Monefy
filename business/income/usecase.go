package income

type IncomeUsecase struct {
	Repo Repository
}

func NewIncomeUsecase(repository Repository) Usecase {
	return &IncomeUsecase{
		Repo: repository,
	}
}

func (usecase *IncomeUsecase) Create(userId uint, income *IncomeDomain) (IncomeDomain, error) {
	income.UserID = userId
	result, err := usecase.Repo.Create(income)
	if err != nil {
		return IncomeDomain{}, err
	}
	return result, nil
}
