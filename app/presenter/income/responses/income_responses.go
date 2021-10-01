package responses

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/income"
)

type Income struct {
	ID          uint      `json:"id"`
	Total       int       `json:"total"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain income.IncomeDomain) Income {
	return Income{
		ID:          domain.ID,
		Description: domain.Description,
		Total:       domain.Total,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromArrayDomain(domain []income.IncomeDomain) []Income {
	var income []Income
	for _, value := range domain {
		income = append(income, Income{
			ID:          value.ID,
			Description: value.Description,
			Total:       value.Total,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}
	return income
}
