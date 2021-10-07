package investplans

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/Rahmanwghazi/Monefy/business"
	"github.com/Rahmanwghazi/Monefy/business/products"
)

type InvestPlanUsecase struct {
	Repo    Repository
	Product products.Repository
}

func NewInvestPlanUsecase(repository Repository, product products.Repository) Usecase {
	return &InvestPlanUsecase{
		Repo:    repository,
		Product: product,
	}
}

func (usecase *InvestPlanUsecase) Create(idProduct string, domain *InvestPlanDomain) (InvestPlanDomain, error) {

	product, err := usecase.Product.GetProductByID(idProduct)
	if strings.TrimSpace(idProduct) != "" {
		if err != nil {
			log.Default().Printf("%+v", err)
		}
		jsonMarshal, err := json.Marshal(product)
		if err != nil {
			log.Default().Printf("%+v", err)
		}
		var desc products.ProductDomain
		json.Unmarshal(jsonMarshal, &desc)
		domain.Description = "ID: " + strconv.Itoa(desc.Data.Product.ID) + "| Name: " + desc.Data.Product.Name + " | Management: " + desc.Data.Product.Management + " | Custodian: " + desc.Data.Product.Custodian + " | Version: " + desc.Data.Version
	}

	result, err := usecase.Repo.Create(domain)
	if err != nil {
		return InvestPlanDomain{}, err
	}
	return result, nil
}

func (usecase *InvestPlanUsecase) GetPlans(domain *InvestPlanDomain) ([]InvestPlanDomain, error) {
	result, err := usecase.Repo.GetPlans(domain)
	if err != nil {
		return []InvestPlanDomain{}, err
	}
	return result, nil
}

func (usecase *InvestPlanUsecase) GetUnfinishedPlans(domain *InvestPlanDomain) ([]InvestPlanDomain, error) {
	result, err := usecase.Repo.GetUnfinishedPlans(domain)
	if err != nil {
		return []InvestPlanDomain{}, err
	}
	return result, nil
}

func (usecase *InvestPlanUsecase) GetfinishedPlans(domain *InvestPlanDomain) ([]InvestPlanDomain, error) {
	result, err := usecase.Repo.GetfinishedPlans(domain)
	if err != nil {
		return []InvestPlanDomain{}, err
	}
	return result, nil
}

func (usecase *InvestPlanUsecase) GetPlanById(domain *InvestPlanDomain, id uint) (InvestPlanDomain, error) {
	result, err := usecase.Repo.GetPlanById(domain, id)
	if err != nil {
		return InvestPlanDomain{}, err
	}
	return result, nil
}

func (usecase *InvestPlanUsecase) EditPlan(domain *InvestPlanDomain, id uint) (InvestPlanDomain, error) {
	_, getErr := usecase.GetPlanById(domain, id)
	if getErr != nil {
		return InvestPlanDomain{}, getErr
	}

	idProduct := strconv.Itoa(domain.ProductID)
	product, err := usecase.Product.GetProductByID(idProduct)
	if strings.TrimSpace(idProduct) != "" {
		if err != nil {
			log.Default().Printf("%+v", err)
		}
		jsonMarshal, err := json.Marshal(product)
		if err != nil {
			log.Default().Printf("%+v", err)
		}
		var desc products.ProductDomain
		json.Unmarshal(jsonMarshal, &desc)
		domain.ID = id
		domain.Description = "ID: " + strconv.Itoa(desc.Data.Product.ID) + "| Name: " + desc.Data.Product.Name + " | Management: " + desc.Data.Product.Management + " | Custodian: " + desc.Data.Product.Custodian + " | Version: " + desc.Data.Version
	}

	result, err := usecase.Repo.EditPlan(domain, id)
	if err != nil {
		return InvestPlanDomain{}, err
	}
	return result, nil
}

func (usecase *InvestPlanUsecase) DeletePlan(domain *InvestPlanDomain, id uint) (string, error) {
	result, err := usecase.Repo.DeletePlan(domain, id)
	if err != nil {
		return business.ErrorInternal.Error(), err
	}
	return result, nil
}
