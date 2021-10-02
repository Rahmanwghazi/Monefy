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
