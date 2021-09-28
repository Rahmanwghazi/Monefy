package income

import (
	"context"

	"gorm.io/gorm"
)

type IncomeDomain struct {
	gorm.Model
	UserID      int
	Total       int
	Description string
}

type Usecase interface {
	Create(context context.Context, domain IncomeDomain) (IncomeDomain, error)
}

type Repository interface {
	Create(context context.Context, domain IncomeDomain) (IncomeDomain, error)
}
