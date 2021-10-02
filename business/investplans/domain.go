package investplans

import "time"

type InvestPlanDomain struct {
	ID          uint
	UserID      uint
	ProductID   int
	Total       int
	DueDate     time.Time
	Description string
	PlanStatus  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	Create(idProduct string, domain InvestPlanDomain) (InvestPlanDomain, error)
	GetPlans(domain InvestPlanDomain) ([]InvestPlanDomain, error)
	EditPlan(domain InvestPlanDomain, id uint) (InvestPlanDomain, error)
}

type Repository interface {
	Create(domain InvestPlanDomain) (InvestPlanDomain, error)
	GetPlans(domain InvestPlanDomain) ([]InvestPlanDomain, error)
	EditPlan(domain InvestPlanDomain, id uint) (InvestPlanDomain, error)
}
