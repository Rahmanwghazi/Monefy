package expenses_test

import (
	"testing"
	"time"

	"github.com/Rahmanwghazi/Monefy/business/expenses"
	_expenseMock "github.com/Rahmanwghazi/Monefy/business/expenses/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockExpenseRepository _expenseMock.Repository
	expenseUsecase        expenses.Usecase

	domainTest expenses.ExpenseDomain
)

func TestMain(m *testing.M) {
	expenseUsecase = expenses.NewExpenseUsecase(&mockExpenseRepository)
	domainTest = expenses.ExpenseDomain{
		ID:          1,
		UserID:      1,
		Total:       500000,
		Description: "Beli mouse",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	m.Run()
}

func TestCreateExpense(t *testing.T) {
	t.Run("Valid test", func(t *testing.T) {
		mockExpenseRepository.On("CreateExpense", mock.Anything).Return(domainTest, nil).Once()

		expense := expenses.ExpenseDomain{
			ID:          1,
			UserID:      1,
			Total:       500000,
			Description: "Beli mouse",
		}

		result, err := expenseUsecase.CreateExpense(&expense)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})
}

//coverage 19%
