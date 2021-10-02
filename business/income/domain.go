package income

import "time"

type IncomeDomain struct {
	ID          uint
	UserID      uint
	Total       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	CreateIncome(domain IncomeDomain) (IncomeDomain, error)
	GetIncome(domain IncomeDomain) ([]IncomeDomain, error)
	EditIncome(domain IncomeDomain, id uint) (IncomeDomain, error)
}

type Repository interface {
	CreateIncome(domain IncomeDomain) (IncomeDomain, error)
	GetIncome(domain IncomeDomain) ([]IncomeDomain, error)
	EditIncome(domain IncomeDomain, id uint) (IncomeDomain, error)
}
