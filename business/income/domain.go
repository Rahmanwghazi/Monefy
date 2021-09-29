package income

import (
	"gorm.io/gorm"
)

type IncomeDomain struct {
	gorm.Model
	UserID      uint
	Total       int
	Description string
}

type Usecase interface {
	Create(userIdomain *IncomeDomain) (IncomeDomain, error)
}

type Repository interface {
	Create(domain *IncomeDomain) (IncomeDomain, error)
}
