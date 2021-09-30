package requests

import "github.com/Rahmanwghazi/Monefy/business/outcome"

type CreateOutcome struct {
	UserID      uint   `json:"user_id"`
	Total       int    `json:"total"`
	Description string `json:"description"`
}

func (createOutcome *CreateOutcome) ToDomain() outcome.OutcomeDomain {
	return outcome.OutcomeDomain{
		UserID:      createOutcome.UserID,
		Total:       createOutcome.Total,
		Description: createOutcome.Description,
	}
}
