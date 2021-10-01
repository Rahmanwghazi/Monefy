package expenses

import (
	"github.com/Rahmanwghazi/Monefy/business/expenses"
	"gorm.io/gorm"
)

type mysqlExpenseRepository struct {
	Connection *gorm.DB
}

func NewMysqlExpenseRepository(connection *gorm.DB) expenses.Repository {
	return &mysqlExpenseRepository{
		Connection: connection,
	}
}

func (rep *mysqlExpenseRepository) Create(domain expenses.ExpenseDomain) (expenses.ExpenseDomain, error) {
	expenseData := FromDomain(domain)
	result := rep.Connection.Create(&expenseData)

	if result.Error != nil {
		return expenses.ExpenseDomain{}, result.Error
	}

	return expenseData.ToDomain(), nil
}

func (rep *mysqlExpenseRepository) GetExpense(domain expenses.ExpenseDomain) ([]expenses.ExpenseDomain, error) {
	var expenseData []Expense
	result := rep.Connection.Find(&expenseData, "user_id = ?", domain.UserID)

	if result.Error != nil {
		return []expenses.ExpenseDomain{}, result.Error
	}

	return ToArrayDomain(expenseData, domain), nil
}
