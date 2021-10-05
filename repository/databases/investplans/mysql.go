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

func (rep *mysqlInvestPlanRepository) GetProducts(domain investplans.InvestPlanDomain) ([]investplans.InvestPlanDomain, error) {
	var product []InvestPlan
	result := rep.Connection.Find(&product)

	if result.Error != nil {
		return []investplans.InvestPlanDomain{}, result.Error
	}

	return ToArrayDomain(product, domain), nil
}

func (rep *mysqlInvestPlanRepository) GetPlans(domain investplans.InvestPlanDomain) ([]investplans.InvestPlanDomain, error) {
	var planData []InvestPlan
	result := rep.Connection.Find(&planData, "user_id = ?", domain.UserID)

	if result.Error != nil {
		return []investplans.InvestPlanDomain{}, result.Error
	}

	return ToArrayDomain(planData, domain), nil
}

func (rep *mysqlInvestPlanRepository) GetPlanById(domain investplans.InvestPlanDomain, id uint) (investplans.InvestPlanDomain, error) {
	var planData InvestPlan
	result := rep.Connection.First(&planData, "user_id = ? AND ID = ? ", domain.UserID, id)

	if result.Error != nil {
		return investplans.InvestPlanDomain{}, result.Error
	}

	return planData.ToDomain(), nil
}

func (rep *mysqlInvestPlanRepository) EditPlan(domain investplans.InvestPlanDomain, id uint) (investplans.InvestPlanDomain, error) {
	planData := FromDomain(domain)

	result := rep.Connection.Where("ID = ? AND user_id = ?", id, domain.UserID).Updates(&planData)

	if result.Error != nil {
		return investplans.InvestPlanDomain{}, result.Error
	}
	return planData.ToDomain(), nil
}

func (rep *mysqlInvestPlanRepository) DeletePlan(domain investplans.InvestPlanDomain, id uint) (string, error) {
	var investPlanData InvestPlan
	result := rep.Connection.Delete(&investPlanData, "user_id = ? AND id = ?", domain.UserID, id)

	if result.Error != nil {
		return "Failed to delete", result.Error
	}

	return "Deleted", nil
}
