package responses

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/investplans"
)

type InvestPlan struct {
	ID          uint      `json:"id"`
	ProductID   int       `json:"product_id"`
	Total       int       `json:"total"`
	DueDate     time.Time `json:"due_date"`
	Description string    `json:"description"`
	PlanStatus  int       `json:"plan_status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain investplans.InvestPlanDomain) InvestPlan {
	return InvestPlan{
		ID:          domain.ID,
		ProductID:   domain.ProductID,
		Total:       domain.Total,
		DueDate:     domain.DueDate,
		Description: domain.Description,
		PlanStatus:  domain.PlanStatus,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromArrayDomain(domain []investplans.InvestPlanDomain) []InvestPlan {
	var investPlan []InvestPlan
	for _, value := range domain {
		investPlan = append(investPlan, InvestPlan{
			ID:          value.ID,
			ProductID:   value.ProductID,
			Total:       value.Total,
			DueDate:     value.DueDate,
			Description: value.Description,
			PlanStatus:  value.PlanStatus,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}
	return investPlan
}
