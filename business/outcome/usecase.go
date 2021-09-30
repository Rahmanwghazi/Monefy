package outcome

type OutcomeUsecase struct {
	Repo Repository
}

func NewOutcomeUsecase(repository Repository) Usecase {
	return &OutcomeUsecase{
		Repo: repository,
	}
}

func (usecase *OutcomeUsecase) Create(outcome *OutcomeDomain) (OutcomeDomain, error) {

	result, err := usecase.Repo.Create(outcome)
	if err != nil {
		return OutcomeDomain{}, err
	}
	return result, nil
}
