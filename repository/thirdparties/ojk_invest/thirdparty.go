package ojkinvest

import (
	"encoding/json"
	"net/http"

	"github.com/Rahmanwghazi/Monefy/business/products"
)

type OJKAPI struct {
	httpClient http.Client
}

func NewOJKAPI() products.Repository {
	return &OJKAPI{
		httpClient: http.Client{},
	}
}

func (product *OJKAPI) GetProductByID(ID string) (products.ProductDomain, error) {
	request, _ := http.NewRequest("GET", "https://ojk-invest-api.vercel.app/api/products/"+ID, nil)
	response, err := product.httpClient.Do(request)
	if err != nil {
		return products.ProductDomain{}, err
	}

	defer response.Body.Close()

	data := Products{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return products.ProductDomain{}, err
	}

	return data.toDomain(), nil
}

func (product *OJKAPI) SearchProduct(domain products.ProductDomain) ([]products.ProductDomain, error) {
	request, _ := http.NewRequest("GET", "https://ojk-invest-api.vercel.app/api/products/", nil)
	response, err := product.httpClient.Do(request)
	if err != nil {
		return []products.ProductDomain{}, err
	}

	defer response.Body.Close()

	var data []Products
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return []products.ProductDomain{}, err
	}

	return ToArrayDomain(data, domain), nil
}
