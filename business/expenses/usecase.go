package expenses

import "github.com/Rahmanwghazi/Monefy/business"

type ExpenseUsecase struct {
	Repo Repository
}

func NewExpenseUsecase(repository Repository) Usecase {
	return &ExpenseUsecase{
		Repo: repository,
	}
}

func (usecase *ExpenseUsecase) CreateExpense(expense *ExpenseDomain) (ExpenseDomain, error) {
	result, err := usecase.Repo.CreateExpense(expense)
	if err != nil {
		return ExpenseDomain{}, err
	}
	return result, nil
}

func (usecase *ExpenseUsecase) GetExpenses(expense *ExpenseDomain) ([]ExpenseDomain, error) {
	result, err := usecase.Repo.GetExpenses(expense)
	if err != nil {
		return []ExpenseDomain{}, err
	}
	return result, nil
}

func (usecase *ExpenseUsecase) GetExpenseById(expense *ExpenseDomain, id uint) (ExpenseDomain, error) {
	result, err := usecase.Repo.GetExpenseById(expense, id)
	if err != nil {
		return ExpenseDomain{}, err
	}
	return result, nil
}

func (usecase *ExpenseUsecase) EditExpense(expense *ExpenseDomain, id uint) (ExpenseDomain, error) {
	result, err := usecase.Repo.EditExpense(expense, id)
	if err != nil {
		return ExpenseDomain{}, err
	}
	return result, nil
}

func (usecase *ExpenseUsecase) DeleteExpense(expense *ExpenseDomain, id uint) (string, error) {
	result, err := usecase.Repo.DeleteExpense(expense, id)
	if err != nil {
		return business.ErrorInternal.Error(), err
	}
	return result, nil
}
