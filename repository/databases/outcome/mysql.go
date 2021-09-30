package outcome

import (
	"github.com/Rahmanwghazi/Monefy/business/outcome"
	"gorm.io/gorm"
)

type mysqlOutcomeRepository struct {
	Connection *gorm.DB
}

func NewMysqlOutcomeRepository(connection *gorm.DB) outcome.Repository {
	return &mysqlOutcomeRepository{
		Connection: connection,
	}
}

func (rep *mysqlOutcomeRepository) Create(domain *outcome.OutcomeDomain) (outcome.OutcomeDomain, error) {
	outcomeData := FromDomain(*domain)
	result := rep.Connection.Create(&outcomeData)

	if result.Error != nil {
		return outcome.OutcomeDomain{}, result.Error
	}

	return outcomeData.ToDomain(), nil
}
