package requests

import (
	"github.com/Rahmanwghazi/Monefy/business/expenses"
)

type Expense struct {
	UserID      uint
	Total       int    `json:"total"`
	Description string `json:"description"`
}

func (expense *Expense) ToDomain() *expenses.ExpenseDomain {
	return &expenses.ExpenseDomain{
		UserID:      expense.UserID,
		Total:       expense.Total,
		Description: expense.Description,
	}
}
