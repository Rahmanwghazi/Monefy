package income

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/income"

	"gorm.io/gorm"
)

type Income struct {
	ID          uint
	UserID      uint
	Total       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (incomeData *Income) ToDomain() income.IncomeDomain {
	return income.IncomeDomain{
		ID:          incomeData.ID,
		UserID:      incomeData.UserID,
		Total:       incomeData.Total,
		Description: incomeData.Description,
		CreatedAt:   incomeData.CreatedAt,
		UpdatedAt:   incomeData.UpdatedAt,
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
		ID:          domain.ID,
		UserID:      domain.UserID,
		Total:       domain.Total,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
