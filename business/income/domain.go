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
	Create(domain IncomeDomain) (IncomeDomain, error)
	GetIncome(domain IncomeDomain) ([]IncomeDomain, error)
}

type Repository interface {
	Create(domain IncomeDomain) (IncomeDomain, error)
	GetIncome(domain IncomeDomain) ([]IncomeDomain, error)
}
