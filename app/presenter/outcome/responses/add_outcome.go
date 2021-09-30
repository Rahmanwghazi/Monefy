package responses

import "github.com/Rahmanwghazi/Monefy/business/outcome"

type CreateOutcome struct {
	Total       int    `json:"total"`
	Description string `json:"description"`
}

func FromDomain(domain outcome.OutcomeDomain) CreateOutcome {
	return CreateOutcome{
		Description: domain.Description,
		Total:       domain.Total,
	}
}
