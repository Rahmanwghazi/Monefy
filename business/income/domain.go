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
	GetIncomeById(domain IncomeDomain, id uint) (IncomeDomain, error)
	EditIncome(domain IncomeDomain, id uint) (IncomeDomain, error)
	DeleteIncome(domain IncomeDomain, id uint) (string, error)
}

type Repository interface {
	CreateIncome(domain IncomeDomain) (IncomeDomain, error)
	GetIncome(domain IncomeDomain) ([]IncomeDomain, error)
	GetIncomeById(domain IncomeDomain, id uint) (IncomeDomain, error)
	EditIncome(domain IncomeDomain, id uint) (IncomeDomain, error)
	DeleteIncome(domain IncomeDomain, id uint) (string, error)
}
