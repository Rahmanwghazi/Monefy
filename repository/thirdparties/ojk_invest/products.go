package ojkinvest

import (
	"github.com/Rahmanwghazi/Monefy/business/products"
)

type Products struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Management string `json:"management"`
	Custodian  string `json:"custodian"`
	Type       string `json:"type"`
}

func (productResponse *Products) toDomain() products.ProductDomain {
	return products.ProductDomain{
		ID:         productResponse.ID,
		Name:       productResponse.Name,
		Maangement: productResponse.Management,
		Custodian:  productResponse.Custodian,
		Type:       productResponse.Type,
	}
}
