package investplans

import (
	"encoding/json"
	"log"
	"strings"

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

func (usecase *InvestPlanUsecase) Create(idProduct string, investplanDomain InvestPlanDomain) (InvestPlanDomain, error) {

	product, err := usecase.Product.GetProductByID(idProduct)
	if strings.TrimSpace(idProduct) != "" {
		if err != nil {
			log.Default().Printf("%+v", err)
		}
		jsonMarshal, err := json.Marshal(product)
		if err != nil {
			log.Default().Printf("%+v", err)
		}
		investplanDomain.Description = string(jsonMarshal)
	}

	result, err := usecase.Repo.Create(investplanDomain)
	if err != nil {
		return InvestPlanDomain{}, err
	}
	return result, nil
}
