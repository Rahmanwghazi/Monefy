package users

import (
	"context"
	"time"
	//"gorm.io/gorm"
)

type Domain struct {
	//gorm.Model
	Username string
	Email    string
	Password string
	FullName string
	DoB      time.Time
	//	Token    string
}

type Usecase interface {
	Signup(context context.Context, domain Domain) (Domain, error)
	Signin(context context.Context, username string, password string) (Domain, error)
}

type Repository interface {
	Signup(context context.Context, domain Domain) (Domain, error)
	Signin(context context.Context, username string, password string) (Domain, error)
}
