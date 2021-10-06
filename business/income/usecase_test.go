package income_test

import (
	"testing"
	"time"

	"github.com/Rahmanwghazi/Monefy/business/income"
	_incomeMock "github.com/Rahmanwghazi/Monefy/business/income/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockincomeRepository _incomeMock.Repository
	incomeUsecase        income.Usecase

	domainTest      income.IncomeDomain
	domainTestArray []income.IncomeDomain
)

func TestMain(m *testing.M) {
	incomeUsecase = income.NewIncomeUsecase(&mockincomeRepository)
	domainTest = income.IncomeDomain{
		ID:          1,
		UserID:      1,
		Total:       500000,
		Description: "Gaji oktober",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	domainTestArray = []income.IncomeDomain{
		{
			ID:          1,
			UserID:      1,
			Total:       500000,
			Description: "Gaji oktober",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			UserID:      2,
			Total:       500000,
			Description: "Gaji oktober",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	m.Run()
}

func TestCreateincome(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockincomeRepository.On("CreateIncome", mock.Anything).Return(domainTest, nil).Once()

		income := income.IncomeDomain{
			ID:          1,
			UserID:      1,
			Total:       500000,
			Description: "Gaji oktober",
		}

		result, err := incomeUsecase.CreateIncome(&income)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})
	t.Run("Test case 2 - Invalid (Unauthorized)", func(t *testing.T) {
		mockincomeRepository.On("CreateIncome", mock.Anything).Return(domainTest, assert.AnError).Once()

		income := income.IncomeDomain{
			ID:          1,
			UserID:      2,
			Total:       500000,
			Description: "Gaji oktober",
		}

		result, err := incomeUsecase.CreateIncome(&income)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, result)
	})
}

func TestGetincome(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockincomeRepository.On("GetIncome", mock.Anything).Return(domainTestArray, nil).Once()

		income := income.IncomeDomain{
			UserID: 1,
		}

		result, err := incomeUsecase.GetIncome(&income)
		assert.Nil(t, err)
		assert.Equal(t, domainTestArray, result)
	})

	t.Run("Test case 2 - Invalid (Empty data)", func(t *testing.T) {
		mockincomeRepository.On("GetIncome", mock.Anything).Return(domainTestArray, assert.AnError).Once()

		income := income.IncomeDomain{
			UserID: 1,
		}

		result, err := incomeUsecase.GetIncome(&income)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTestArray, result)
	})
}

func TestGetincomeById(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockincomeRepository.On("GetIncomeById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, nil).Once()
		id := uint(1)
		income := income.IncomeDomain{
			UserID: 1,
		}

		result, err := incomeUsecase.GetIncomeById(&income, id)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})

	t.Run("Test case 2 - Invalid (Empty data)", func(t *testing.T) {
		mockincomeRepository.On("GetIncomeById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, assert.AnError).Once()
		id := uint(3)
		income := income.IncomeDomain{
			UserID: 1,
		}

		result, err := incomeUsecase.GetIncomeById(&income, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, result)
	})
}

func TestEditincome(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockincomeRepository.On("EditIncome", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, nil).Once()
		id := uint(1)
		income := income.IncomeDomain{
			ID:          1,
			UserID:      1,
			Total:       500000,
			Description: "Gaji oktober",
		}

		result, err := incomeUsecase.EditIncome(&income, id)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})

	t.Run("Test case 2 - Invalid (Empty field)", func(t *testing.T) {
		mockincomeRepository.On("EditIncome", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, assert.AnError).Once()
		id := uint(1)
		income := income.IncomeDomain{
			ID:     1,
			UserID: 1,
		}

		result, err := incomeUsecase.EditIncome(&income, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, result)
	})
}

func TestDeleteincome(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockincomeRepository.On("DeleteIncome", mock.Anything, mock.AnythingOfType("uint")).Return("Deleted", nil).Once()
		id := uint(1)
		income := income.IncomeDomain{
			ID:     1,
			UserID: 1,
		}

		result, err := incomeUsecase.DeleteIncome(&income, id)
		assert.Nil(t, err)
		assert.Equal(t, "Deleted", result)
	})

	t.Run("Test case 2 - Invalid (Unauthorized)", func(t *testing.T) {
		mockincomeRepository.On("DeleteIncome", mock.Anything, mock.AnythingOfType("uint")).Return("Deleted", assert.AnError).Once()
		id := uint(1)
		income := income.IncomeDomain{
			ID:     1,
			UserID: 2,
		}

		result, err := incomeUsecase.DeleteIncome(&income, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, "Deleted", result)
	})
}

//coverage 100.0%
