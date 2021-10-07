package responses

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/expenses"
)

type Expense struct {
	ID          uint      `json:"id"`
	Total       int       `json:"total"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain expenses.ExpenseDomain) Expense {
	return Expense{
		ID:          domain.ID,
		Description: domain.Description,
		Total:       domain.Total,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromArrayDomain(domain []expenses.ExpenseDomain) []Expense {
	var expense []Expense
	for _, value := range domain {
		expense = append(expense, Expense{
			ID:          value.ID,
			Description: value.Description,
			Total:       value.Total,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}
	return expense
}
