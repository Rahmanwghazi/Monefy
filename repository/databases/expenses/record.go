package expenses

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/expenses"

	"gorm.io/gorm"
)

type Expense struct {
	ID          uint
	UserID      uint
	Total       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (expenseData *Expense) ToDomain() expenses.ExpenseDomain {
	return expenses.ExpenseDomain{
		ID:          expenseData.ID,
		UserID:      expenseData.UserID,
		Total:       expenseData.Total,
		Description: expenseData.Description,
		CreatedAt:   expenseData.CreatedAt,
		UpdatedAt:   expenseData.UpdatedAt,
	}
}

func ToArrayDomain(expenseData []Expense, domain expenses.ExpenseDomain) []expenses.ExpenseDomain {
	result := []expenses.ExpenseDomain{}
	for _, expense := range expenseData {
		result = append(result, expense.ToDomain())
	}
	return result
}

func FromDomain(domain expenses.ExpenseDomain) *Expense {
	return &Expense{
		ID:          domain.ID,
		UserID:      domain.UserID,
		Total:       domain.Total,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
