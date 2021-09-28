package income

import (
	"github.com/Rahmanwghazi/Monefy/business/income"
	"gorm.io/gorm"
)

type Income struct {
	gorm.Model
	UserID      uint
	Total       int
	Description string
}

func (incomeData *Income) ToDomain() income.IncomeDomain {
	return income.IncomeDomain{
		Model:       gorm.Model{},
		Total:       incomeData.Total,
		Description: incomeData.Description,
	}
}

func FromDomain(domain income.IncomeDomain) Income {
	return Income{
		Model:       gorm.Model{},
		UserID:      uint(domain.UserID),
		Total:       domain.Total,
		Description: domain.Description,
	}
}
