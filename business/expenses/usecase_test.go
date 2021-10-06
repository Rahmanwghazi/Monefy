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

	domainTest      expenses.ExpenseDomain
	domainTestArray []expenses.ExpenseDomain
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
	domainTestArray = []expenses.ExpenseDomain{
		{
			ID:          1,
			UserID:      1,
			Total:       500000,
			Description: "Beli mouse",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	m.Run()
}

func TestCreateExpense(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
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

	t.Run("Test case 2 - Invalid (Internal server error)", func(t *testing.T) {
		mockExpenseRepository.On("CreateExpense", mock.Anything).Return(domainTest, assert.AnError).Once()

		expense := expenses.ExpenseDomain{
			ID:          1,
			UserID:      1,
			Total:       500000,
			Description: "Beli mouse",
		}

		result, err := expenseUsecase.CreateExpense(&expense)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, result)
	})
}

func TestGetExpenses(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockExpenseRepository.On("GetExpenses", mock.Anything).Return(domainTestArray, nil).Once()

		expense := expenses.ExpenseDomain{
			UserID: 1,
		}

		result, err := expenseUsecase.GetExpenses(&expense)
		assert.Nil(t, err)
		assert.Equal(t, domainTestArray, result)
	})

	t.Run("Test case 2 - Invalid (Empty data)", func(t *testing.T) {
		mockExpenseRepository.On("GetExpenses", mock.Anything).Return(domainTestArray, assert.AnError).Once()

		expense := expenses.ExpenseDomain{
			UserID: 1,
		}

		result, err := expenseUsecase.GetExpenses(&expense)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTestArray, result)
	})
}

func TestGetExpensesById(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockExpenseRepository.On("GetExpenseById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, nil).Once()
		id := uint(1)
		expense := expenses.ExpenseDomain{
			UserID: 1,
		}

		result, err := expenseUsecase.GetExpenseById(&expense, id)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})

	t.Run("Test case 2 - Invalid (Empty data)", func(t *testing.T) {
		mockExpenseRepository.On("GetExpenseById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, assert.AnError).Once()
		id := uint(3)
		expense := expenses.ExpenseDomain{
			UserID: 1,
		}

		result, err := expenseUsecase.GetExpenseById(&expense, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, result)
	})
}

func TestEditExpense(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockExpenseRepository.On("GetExpenseById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, nil).Once()
		mockExpenseRepository.On("EditExpense", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, nil).Once()
		id := uint(1)
		expense := expenses.ExpenseDomain{
			ID:          1,
			UserID:      1,
			Total:       500000,
			Description: "Beli keyboard",
		}

		result, err := expenseUsecase.EditExpense(&expense, id)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})

	t.Run("Test case 2 - Invalid (Empty field)", func(t *testing.T) {
		mockExpenseRepository.On("GetExpenseById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, assert.AnError).Once()
		mockExpenseRepository.On("EditExpense", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, assert.AnError).Once()
		id := uint(1)
		expense := expenses.ExpenseDomain{
			ID:          1,
			UserID:      2,
			Total:       3500000,
			Description: "Beli HP",
		}

		result, err := expenseUsecase.EditExpense(&expense, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, result)
	})
}

func TestDeleteExpense(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockExpenseRepository.On("DeleteExpense", mock.Anything, mock.AnythingOfType("uint")).Return("Deleted", nil).Once()
		id := uint(1)
		expense := expenses.ExpenseDomain{
			ID: 1,
		}

		result, err := expenseUsecase.DeleteExpense(&expense, id)
		assert.Nil(t, err)
		assert.Equal(t, "Deleted", result)
	})

	t.Run("Test case 2 - Invalid (Unauthorized)", func(t *testing.T) {
		mockExpenseRepository.On("DeleteExpense", mock.Anything, mock.AnythingOfType("uint")).Return("Deleted", assert.AnError).Once()
		id := uint(1)
		expense := expenses.ExpenseDomain{
			ID: 2,
		}

		result, err := expenseUsecase.DeleteExpense(&expense, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, "Deleted", result)
	})
}

//coverage 96.0%
