package requests

import "github.com/Rahmanwghazi/Monefy/business/income"

type CreateIncome struct {
	Total       int    `json:"total"`
	Description string `json:"description"`
}

func (createIncome *CreateIncome) ToDomain() income.IncomeDomain {
	return income.IncomeDomain{
		Total:       createIncome.Total,
		Description: createIncome.Description,
	}
}
