package expenses

import "time"

type ExpenseDomain struct {
	ID          uint
	UserID      uint
	Total       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	CreateExpense(domain ExpenseDomain) (ExpenseDomain, error)
	GetExpenses(domain ExpenseDomain) ([]ExpenseDomain, error)
	GetExpenseById(domain ExpenseDomain, id uint) (ExpenseDomain, error)
	EditExpense(domain ExpenseDomain, id uint) (ExpenseDomain, error)
	DeleteExpense(domain ExpenseDomain, id uint) (string, error)
}

type Repository interface {
	CreateExpense(domain ExpenseDomain) (ExpenseDomain, error)
	GetExpenses(domain ExpenseDomain) ([]ExpenseDomain, error)
	GetExpenseById(domain ExpenseDomain, id uint) (ExpenseDomain, error)
	EditExpense(domain ExpenseDomain, id uint) (ExpenseDomain, error)
	DeleteExpense(domain ExpenseDomain, id uint) (string, error)
}
