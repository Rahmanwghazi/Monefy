package outcome

import (
	"github.com/Rahmanwghazi/Monefy/business/outcome"

	"gorm.io/gorm"
)

type Outcome struct {
	gorm.Model
	UserID      uint
	Total       int
	Description string
}

func (outcomeData *Outcome) ToDomain() outcome.OutcomeDomain {
	return outcome.OutcomeDomain{
		Model:       gorm.Model{},
		UserID:      outcomeData.UserID,
		Total:       outcomeData.Total,
		Description: outcomeData.Description,
	}
}

func FromDomain(domain outcome.OutcomeDomain) *Outcome {
	return &Outcome{
		Model:       gorm.Model{},
		UserID:      domain.UserID,
		Total:       domain.Total,
		Description: domain.Description,
	}
}
