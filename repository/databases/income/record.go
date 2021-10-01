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
		UserID:      incomeData.UserID,
		Total:       incomeData.Total,
		Description: incomeData.Description,
	}
}

func ToArrayDomain(incomeData []Income, domain income.IncomeDomain) []income.IncomeDomain {
	result := []income.IncomeDomain{}
	for _, income := range incomeData {
		result = append(result, income.ToDomain())
	}
	return result
}

func FromDomain(domain income.IncomeDomain) *Income {
	return &Income{
		Model:       gorm.Model{},
		UserID:      domain.UserID,
		Total:       domain.Total,
		Description: domain.Description,
	}
}
