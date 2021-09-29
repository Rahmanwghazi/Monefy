package requests

import "github.com/Rahmanwghazi/Monefy/business/income"

type CreateIncome struct {
	UserID      uint   `json:"user_id"`
	Total       int    `json:"total"`
	Description string `json:"description"`
}

func (createIncome *CreateIncome) ToDomain() income.IncomeDomain {
	return income.IncomeDomain{
		UserID:      createIncome.UserID,
		Total:       createIncome.Total,
		Description: createIncome.Description,
	}
}
