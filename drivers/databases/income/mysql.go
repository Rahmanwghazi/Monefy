package income

import (
	"github.com/Rahmanwghazi/Monefy/business/income"
	"gorm.io/gorm"
)

type mysqlIncomeRepository struct {
	Connection *gorm.DB
}

func NewMysqlIncomeRepository(connection *gorm.DB) income.Repository {
	return &mysqlIncomeRepository{
		Connection: connection,
	}
}

func (rep *mysqlIncomeRepository) Create(domain *income.IncomeDomain) (income.IncomeDomain, error) {
	incomeData := Income{}
	incomeData.Total = domain.Total
	incomeData.Description = domain.Description

	result := rep.Connection.Create(&incomeData)

	if result.Error != nil {
		return income.IncomeDomain{}, result.Error
	}

	return incomeData.ToDomain(), nil
}
