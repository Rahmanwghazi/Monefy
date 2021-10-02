package investplans

import (
	"github.com/Rahmanwghazi/Monefy/business/investplans"
	"gorm.io/gorm"
)

type mysqlInvestPlanRepository struct {
	Connection *gorm.DB
}

func NewMysqlnvestPlanRepository(connection *gorm.DB) investplans.Repository {
	return &mysqlInvestPlanRepository{
		Connection: connection,
	}
}

func (rep *mysqlInvestPlanRepository) Create(domain investplans.InvestPlanDomain) (investplans.InvestPlanDomain, error) {
	investPlanData := FromDomain(domain)
	result := rep.Connection.Create(&investPlanData)

	if result.Error != nil {
		return investplans.InvestPlanDomain{}, result.Error
	}

	return investPlanData.ToDomain(), nil
}

func (rep *mysqlInvestPlanRepository) GetPlans(domain investplans.InvestPlanDomain) ([]investplans.InvestPlanDomain, error) {
	var planData []InvestPlan
	result := rep.Connection.Find(&planData, "user_id = ?", domain.UserID)

	if result.Error != nil {
		return []investplans.InvestPlanDomain{}, result.Error
	}

	return ToArrayDomain(planData, domain), nil
}

func (rep *mysqlInvestPlanRepository) EditPlan(domain investplans.InvestPlanDomain, id uint) (investplans.InvestPlanDomain, error) {
	planData := FromDomain(domain)

	result := rep.Connection.Where("ID = ?", id).Updates(&planData)

	if result.Error != nil {
		return investplans.InvestPlanDomain{}, result.Error
	}
	return planData.ToDomain(), nil
}
