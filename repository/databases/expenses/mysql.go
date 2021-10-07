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

func (rep *mysqlExpenseRepository) CreateExpense(domain *expenses.ExpenseDomain) (expenses.ExpenseDomain, error) {
	expenseData := FromDomain(*domain)
	result := rep.Connection.Create(&expenseData)

	if result.Error != nil {
		return expenses.ExpenseDomain{}, result.Error
	}

	return expenseData.ToDomain(), nil
}

func (rep *mysqlExpenseRepository) GetExpenses(domain *expenses.ExpenseDomain) ([]expenses.ExpenseDomain, error) {
	var expenseData []Expense
	result := rep.Connection.Find(&expenseData, "user_id = ?", domain.UserID)

	if result.Error != nil {
		return []expenses.ExpenseDomain{}, result.Error
	}

	return ToArrayDomain(expenseData, *domain), nil
}

func (rep *mysqlExpenseRepository) GetExpenseById(domain *expenses.ExpenseDomain, id uint) (expenses.ExpenseDomain, error) {
	var expenseData Expense
	result := rep.Connection.First(&expenseData, "user_id = ? AND id = ?", domain.UserID, id)

	if result.Error != nil {
		return expenses.ExpenseDomain{}, result.Error
	}

	return expenseData.ToDomain(), nil
}

func (rep *mysqlExpenseRepository) EditExpense(domain *expenses.ExpenseDomain, id uint) (expenses.ExpenseDomain, error) {
	expenseData := FromDomain(*domain)

	result := rep.Connection.Where("ID = ? AND user_id = ?", id, domain.UserID).Updates(&expenseData)

	if result.Error != nil {
		return expenses.ExpenseDomain{}, result.Error
	}

	return expenseData.ToDomain(), nil
}

func (rep *mysqlExpenseRepository) DeleteExpense(domain *expenses.ExpenseDomain, id uint) (string, error) {
	var expenseData Expense
	result := rep.Connection.Delete(&expenseData, "user_id = ? AND id = ?", domain.UserID, id)

	if result.Error != nil {
		return "Failed to delete", result.Error
	}

	return "Deleted", nil
}
