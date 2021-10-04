package responses

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

func FromDomain(domain products.ProductDomain) Products {
	return Products{
		Data: domain.Data,
	}
}

func FromArrayDomain(domain []products.ProductDomain) []Products {
	var product []Products
	for _, value := range domain {
		product = append(product, Products{
			Data: value.Data,
		})
	}
	return product
}
