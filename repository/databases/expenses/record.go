package expenses

import (
	"github.com/Rahmanwghazi/Monefy/business/expenses"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	UserID      uint
	Total       int
	Description string
}

func (expenseData *Expense) ToDomain() expenses.ExpenseDomain {
	return expenses.ExpenseDomain{
		Model:       gorm.Model{},
		UserID:      expenseData.UserID,
		Total:       expenseData.Total,
		Description: expenseData.Description,
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
		Model:       gorm.Model{},
		UserID:      domain.UserID,
		Total:       domain.Total,
		Description: domain.Description,
	}
}
