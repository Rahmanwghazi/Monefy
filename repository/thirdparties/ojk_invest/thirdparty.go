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

func (ipl *OJKAPI) GetProductByID(ID string) (products.ProductDomain, error) {
	req, _ := http.NewRequest("GET", "https://ojk-invest-api.vercel.app/api/products/"+ID, nil)
	resp, err := ipl.httpClient.Do(req)
	if err != nil {
		return products.ProductDomain{}, err
	}

	defer resp.Body.Close()

	data := Products{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return products.ProductDomain{}, err
	}

	return data.toDomain(), nil
}
