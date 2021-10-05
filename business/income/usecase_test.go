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
	mockIncomeRepository _incomeMock.Repository
	incomeUsecase        income.Usecase

	domainTest income.IncomeDomain
)

func TestMain(m *testing.M) {
	incomeUsecase = income.NewIncomeUsecase(&mockIncomeRepository)
	domainTest = income.IncomeDomain{
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
		mockIncomeRepository.On("CreateIncome", mock.Anything).Return(domainTest, nil).Once()

		income := income.IncomeDomain{
			ID:          1,
			UserID:      1,
			Total:       500000,
			Description: "Beli mouse",
		}

		result, err := incomeUsecase.CreateIncome(&income)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})
}

//coverage: 19%
