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

func (rep *mysqlIncomeRepository) CreateIncome(domain *income.IncomeDomain) (income.IncomeDomain, error) {
	incomeData := FromDomain(*domain)
	result := rep.Connection.Create(&incomeData)

	if result.Error != nil {
		return income.IncomeDomain{}, result.Error
	}

	return incomeData.ToDomain(), nil
}

func (rep *mysqlIncomeRepository) GetIncome(domain *income.IncomeDomain) ([]income.IncomeDomain, error) {
	var incomeData []Income
	result := rep.Connection.Find(&incomeData, "user_id = ?", domain.UserID)

	if result.Error != nil {
		return []income.IncomeDomain{}, result.Error
	}

	return ToArrayDomain(incomeData, *domain), nil
}

func (rep *mysqlIncomeRepository) GetIncomeById(domain *income.IncomeDomain, id uint) (income.IncomeDomain, error) {
	var incomeData Income
	result := rep.Connection.First(&incomeData, "user_id = ? AND id = ?", domain.UserID, id)

	if result.Error != nil {
		return income.IncomeDomain{}, result.Error
	}

	return incomeData.ToDomain(), nil
}

func (rep *mysqlIncomeRepository) EditIncome(domain *income.IncomeDomain, id uint) (income.IncomeDomain, error) {
	incomeData := FromDomain(*domain)

	result := rep.Connection.Where("ID = ? AND user_id = ?", id, domain.UserID).Updates(&incomeData)

	if result.Error != nil {
		return income.IncomeDomain{}, result.Error
	}

	return incomeData.ToDomain(), nil
}

func (rep *mysqlIncomeRepository) DeleteIncome(domain *income.IncomeDomain, id uint) (string, error) {
	var incomeData Income
	result := rep.Connection.Delete(&incomeData, "user_id = ? AND id = ?", domain.UserID, id)

	if result.Error != nil {
		return "Failed to delete", result.Error
	}

	return "Deleted", nil
}
