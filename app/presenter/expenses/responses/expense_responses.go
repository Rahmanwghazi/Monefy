package responses

import "github.com/Rahmanwghazi/Monefy/business/expenses"

type Expense struct {
	Total       int    `json:"total"`
	Description string `json:"description"`
}

func FromDomain(domain expenses.ExpenseDomain) Expense {
	return Expense{
		Description: domain.Description,
		Total:       domain.Total,
	}
}

func FromArrayDomain(domain []expenses.ExpenseDomain) []Expense {
	var expense []Expense
	for _, value := range domain {
		expense = append(expense, Expense{
			Description: value.Description,
			Total:       value.Total,
		})
	}
	return expense
}
