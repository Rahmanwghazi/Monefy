package expenses

import (
	"gorm.io/gorm"
)

type ExpenseDomain struct {
	gorm.Model
	UserID      uint
	Total       int
	Description string
}

type Usecase interface {
	Create(domain ExpenseDomain) (ExpenseDomain, error)
	GetExpense(domain ExpenseDomain) ([]ExpenseDomain, error)
}

type Repository interface {
	Create(domain ExpenseDomain) (ExpenseDomain, error)
	GetExpense(domain ExpenseDomain) ([]ExpenseDomain, error)
}
