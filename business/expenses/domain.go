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
	GetExpense(domain ExpenseDomain) ([]ExpenseDomain, error)
	EditExpense(domain ExpenseDomain, id uint) (ExpenseDomain, error)
}

type Repository interface {
	CreateExpense(domain ExpenseDomain) (ExpenseDomain, error)
	GetExpense(domain ExpenseDomain) ([]ExpenseDomain, error)
	EditExpense(domain ExpenseDomain, id uint) (ExpenseDomain, error)
}
