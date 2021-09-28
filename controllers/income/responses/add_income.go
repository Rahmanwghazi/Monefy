package responses

import "github.com/Rahmanwghazi/Monefy/business/income"

type CreateIncome struct {
	Total       int    `json:"total"`
	Description string `json:"description"`
}

func FromDomain(domain income.IncomeDomain) CreateIncome {
	return CreateIncome{
		Description: domain.Description,
		Total:       domain.Total,
	}
}
