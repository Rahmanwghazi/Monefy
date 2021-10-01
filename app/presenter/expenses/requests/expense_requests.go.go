package requests

import "github.com/Rahmanwghazi/Monefy/business/expenses"

type CreateExpense struct {
	UserID      uint   `json:"user_id"`
	Total       int    `json:"total"`
	Description string `json:"description"`
}

func (createExpense *CreateExpense) ToDomain() expenses.ExpenseDomain {
	return expenses.ExpenseDomain{
		UserID:      createExpense.UserID,
		Total:       createExpense.Total,
		Description: createExpense.Description,
	}
}
