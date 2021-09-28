package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserDomain struct {
	gorm.Model
	Username string
	Email    string
	Password string
	FullName string
	DoB      time.Time
	Token    string
}

type Usecase interface {
	Signup(context context.Context, domain UserDomain) (UserDomain, error)
	Signin(context context.Context, username string, password string) (UserDomain, error)
}

type Repository interface {
	Signup(context context.Context, domain UserDomain) (UserDomain, error)
	Signin(context context.Context, username string, password string) (UserDomain, error)
}
