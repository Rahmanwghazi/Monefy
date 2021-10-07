package requests

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/investplans"
)

type InvestPlan struct {
	UserID     uint      `json:"user_id"`
	ProductID  int       `json:"product_id"`
	Total      int       `json:"total"`
	DueDate    time.Time `json:"due_date"`
	PlanStatus int       `json:"plan_status"`
}

func (investPlan *InvestPlan) ToDomain() *investplans.InvestPlanDomain {
	return &investplans.InvestPlanDomain{
		UserID:     investPlan.UserID,
		ProductID:  investPlan.ProductID,
		Total:      investPlan.Total,
		DueDate:    investPlan.DueDate,
		PlanStatus: investPlan.PlanStatus,
	}
}
