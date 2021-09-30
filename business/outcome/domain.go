package outcome

import (
	"gorm.io/gorm"
)

type OutcomeDomain struct {
	gorm.Model
	UserID      uint
	Total       int
	Description string
}

type Usecase interface {
	Create(domain *OutcomeDomain) (OutcomeDomain, error)
}

type Repository interface {
	Create(domain *OutcomeDomain) (OutcomeDomain, error)
}
