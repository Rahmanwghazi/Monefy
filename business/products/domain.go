package products

type ProductDomain struct {
	ID         int
	Name       string
	Maangement string
	Custodian  string
	Type       string
}

type Repository interface {
	GetProductByID(id string) (ProductDomain, error)
}
