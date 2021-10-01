package expenses

type ExpenseUsecase struct {
	Repo Repository
}

func NewExpenseUsecase(repository Repository) Usecase {
	return &ExpenseUsecase{
		Repo: repository,
	}
}

func (usecase *ExpenseUsecase) Create(expense ExpenseDomain) (ExpenseDomain, error) {

	result, err := usecase.Repo.Create(expense)
	if err != nil {
		return ExpenseDomain{}, err
	}
	return result, nil
}

func (usecase *ExpenseUsecase) GetExpense(expense ExpenseDomain) ([]ExpenseDomain, error) {
	result, err := usecase.Repo.GetExpense(expense)
	if err != nil {
		return []ExpenseDomain{}, err
	}
	return result, nil

}
