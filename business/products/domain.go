package products

type ProductDomain struct {
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
}

type Usecase interface {
	SearchProduct(domain ProductDomain) ([]ProductDomain, error)
}

type Repository interface {
	GetProductByID(id string) (ProductDomain, error)
	SearchProduct(domain ProductDomain) ([]ProductDomain, error)
}
