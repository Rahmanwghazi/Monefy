package ojkinvest

import (
	"github.com/Rahmanwghazi/Monefy/business/products"
)

type Products struct {
	Data struct {
		Product struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Management string `json:"management"`
			Custodian  string `json:"custodian"`
			Type       string `json:"type"`
		} `json:"product"`
		Version string `json:"version"`
	} `json:"data"`
	Error interface{} `json:"error"`
}

func (productResponse *Products) toDomain() products.ProductDomain {
	return products.ProductDomain{
		Data: productResponse.Data,
	}
}
