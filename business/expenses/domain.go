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
	Create(domain ExpenseDomain) (ExpenseDomain, error)
	GetExpense(domain ExpenseDomain) ([]ExpenseDomain, error)
}

type Repository interface {
	Create(domain ExpenseDomain) (ExpenseDomain, error)
	GetExpense(domain ExpenseDomain) ([]ExpenseDomain, error)
}
