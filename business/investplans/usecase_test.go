package investplans_test

import (
	"testing"
	"time"

	"github.com/Rahmanwghazi/Monefy/business/investplans"
	_planMock "github.com/Rahmanwghazi/Monefy/business/investplans/mocks"
	"github.com/Rahmanwghazi/Monefy/business/products"
	_productRepositoryMock "github.com/Rahmanwghazi/Monefy/business/products/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockPlanRepository    _planMock.Repository
	domainTest            investplans.InvestPlanDomain
	domainTestArray       []investplans.InvestPlanDomain
	productDomain         products.ProductDomain
	mockProductRepository _productRepositoryMock.Repository
	planUsecase           investplans.Usecase
)

func TestMain(m *testing.M) {
	planUsecase = investplans.NewInvestPlanUsecase(&mockPlanRepository, &mockProductRepository)
	domainTest = investplans.InvestPlanDomain{
		ID:          1,
		UserID:      1,
		ProductID:   15,
		Total:       2500000,
		DueDate:     time.Now(),
		Description: "Sucorinvest",
		PlanStatus:  1,
	}
	domainTestArray = []investplans.InvestPlanDomain{
		{
			ID:          1,
			UserID:      1,
			ProductID:   15,
			Total:       2500000,
			DueDate:     time.Now(),
			Description: "Sucorinvest",
			PlanStatus:  1,
		},
	}
	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockProductRepository.On("GetProductByID", mock.AnythingOfType("string")).Return(productDomain, nil).Once()
		mockPlanRepository.On("Create", mock.Anything).Return(domainTest, nil).Once()

		idProduct := "15"
		input := investplans.InvestPlanDomain{
			ID:          1,
			UserID:      1,
			ProductID:   15,
			Total:       250000,
			DueDate:     time.Now(),
			Description: "Sucorinvest",
			PlanStatus:  1,
		}

		result, err := planUsecase.Create(idProduct, &input)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})

	t.Run("Test case 2 - Invalid (Internal server error)", func(t *testing.T) {
		mockProductRepository.On("GetProductByID", mock.AnythingOfType("string")).Return(productDomain, nil).Once()
		mockPlanRepository.On("Create", mock.Anything).Return(domainTest, assert.AnError).Once()

		idProduct := "15"
		input := investplans.InvestPlanDomain{
			ID:          1,
			UserID:      1,
			ProductID:   15,
			Total:       250000,
			DueDate:     time.Now(),
			Description: "Sucorinvest",
			PlanStatus:  1,
		}

		result, err := planUsecase.Create(idProduct, &input)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, result)
	})
}

func TestGetPlans(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockPlanRepository.On("GetPlans", mock.Anything).Return(domainTestArray, nil).Once()

		input := investplans.InvestPlanDomain{
			UserID: 1,
		}

		result, err := planUsecase.GetPlans(&input)
		assert.Nil(t, err)
		assert.Equal(t, domainTestArray, result)
	})

	t.Run("Test case 2 - Invalid (Empty data)", func(t *testing.T) {
		mockPlanRepository.On("GetPlans", mock.Anything).Return(domainTestArray, assert.AnError).Once()

		input := investplans.InvestPlanDomain{
			UserID: 1,
		}

		result, err := planUsecase.GetPlans(&input)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTestArray, result)
	})
}

func TestGetUnifinishedPlans(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockPlanRepository.On("GetUnfinishedPlans", mock.Anything).Return(domainTestArray, nil).Once()

		input := investplans.InvestPlanDomain{
			UserID: 1,
		}

		result, err := planUsecase.GetUnfinishedPlans(&input)
		assert.Nil(t, err)
		assert.Equal(t, domainTestArray, result)
	})

	t.Run("Test case 2 - Invalid (Empty data))", func(t *testing.T) {
		mockPlanRepository.On("GetUnfinishedPlans", mock.Anything).Return(domainTestArray, assert.AnError).Once()

		input := investplans.InvestPlanDomain{
			UserID: 1,
		}

		result, err := planUsecase.GetUnfinishedPlans(&input)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTestArray, result)
	})
}

func TestGetFinishedPlans(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockPlanRepository.On("GetfinishedPlans", mock.Anything).Return(domainTestArray, nil).Once()

		input := investplans.InvestPlanDomain{
			UserID: 1,
		}

		result, err := planUsecase.GetfinishedPlans(&input)
		assert.Nil(t, err)
		assert.Equal(t, domainTestArray, result)
	})

	t.Run("Test case 2 - Invalid (Empty data))", func(t *testing.T) {
		mockPlanRepository.On("GetfinishedPlans", mock.Anything).Return(domainTestArray, assert.AnError).Once()

		input := investplans.InvestPlanDomain{
			UserID: 1,
		}

		result, err := planUsecase.GetfinishedPlans(&input)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTestArray, result)
	})
}

func TestGetPlanById(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockPlanRepository.On("GetPlanById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, nil).Once()
		id := uint(1)
		input := investplans.InvestPlanDomain{
			ID:     1,
			UserID: 1,
		}

		result, err := planUsecase.GetPlanById(&input, id)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})
	t.Run("Test case 2 - Invalid (Unauthorized)", func(t *testing.T) {
		mockPlanRepository.On("GetPlanById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, assert.AnError).Once()
		id := uint(4)
		input := investplans.InvestPlanDomain{
			ID:     4,
			UserID: 2,
		}

		result, err := planUsecase.GetPlanById(&input, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, result)
	})
}

func TestEditPlan(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockPlanRepository.On("GetPlanById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, nil).Once()
		mockProductRepository.On("GetProductByID", mock.AnythingOfType("string")).Return(productDomain, nil).Once()
		mockPlanRepository.On("EditPlan", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, nil).Once()
		id := uint(1)
		input := investplans.InvestPlanDomain{
			ID:          1,
			UserID:      1,
			ProductID:   15,
			Total:       250000,
			DueDate:     time.Now(),
			Description: "Sucorinvest",
			PlanStatus:  1,
		}

		result, err := planUsecase.EditPlan(&input, id)
		assert.Nil(t, err)
		assert.Equal(t, domainTest, result)
	})

	t.Run("Test case 2 - Invalid (Empty data)", func(t *testing.T) {
		mockPlanRepository.On("GetPlanById", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, nil).Once()
		mockProductRepository.On("GetProductByID", mock.AnythingOfType("string")).Return(productDomain, nil).Once()
		mockPlanRepository.On("EditPlan", mock.Anything, mock.AnythingOfType("uint")).Return(domainTest, assert.AnError).Once()
		id := uint(1)
		input := investplans.InvestPlanDomain{
			ID:          5,
			UserID:      1,
			ProductID:   15,
			Total:       250000,
			DueDate:     time.Now(),
			Description: "Sucorinvest",
			PlanStatus:  1,
		}

		result, err := planUsecase.EditPlan(&input, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, domainTest, result)
	})
}

func TestDeletePlan(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		mockPlanRepository.On("DeletePlan", mock.Anything, mock.AnythingOfType("uint")).Return("Deleted", nil).Once()
		id := uint(1)
		input := investplans.InvestPlanDomain{
			ID:     1,
			UserID: 1,
		}

		result, err := planUsecase.DeletePlan(&input, id)
		assert.Nil(t, err)
		assert.Equal(t, "Deleted", result)
	})

	t.Run("Test case 2 - Invalid (Empty data) ", func(t *testing.T) {
		mockPlanRepository.On("DeletePlan", mock.Anything, mock.AnythingOfType("uint")).Return("Deleted", assert.AnError).Once()
		id := uint(3)
		input := investplans.InvestPlanDomain{
			ID:     3,
			UserID: 1,
		}

		result, err := planUsecase.DeletePlan(&input, id)
		assert.NotNil(t, err)
		assert.NotEqual(t, "Deleted", result)
	})
}

//coverage 90.2%
