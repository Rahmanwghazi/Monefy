package investplans

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/investplans"

	"gorm.io/gorm"
)

type InvestPlan struct {
	ID          uint
	UserID      uint
	ProductID   int
	Total       int
	DueDate     time.Time
	Description string
	PlanStatus  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (investPlanData *InvestPlan) ToDomain() investplans.InvestPlanDomain {
	return investplans.InvestPlanDomain{
		ID:          investPlanData.ID,
		UserID:      investPlanData.UserID,
		ProductID:   investPlanData.ProductID,
		Total:       investPlanData.Total,
		DueDate:     investPlanData.DueDate,
		Description: investPlanData.Description,
		PlanStatus:  investPlanData.PlanStatus,
		CreatedAt:   investPlanData.CreatedAt,
		UpdatedAt:   investPlanData.UpdatedAt,
	}
}

func ToArrayDomain(investPlanData []InvestPlan, domain investplans.InvestPlanDomain) []investplans.InvestPlanDomain {
	result := []investplans.InvestPlanDomain{}
	for _, investPlan := range investPlanData {
		result = append(result, investPlan.ToDomain())
	}
	return result
}

func FromDomain(domain investplans.InvestPlanDomain) *InvestPlan {
	return &InvestPlan{
		ID:          domain.ID,
		UserID:      domain.UserID,
		ProductID:   domain.ProductID,
		Total:       domain.Total,
		DueDate:     domain.DueDate,
		Description: domain.Description,
		PlanStatus:  domain.PlanStatus,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
