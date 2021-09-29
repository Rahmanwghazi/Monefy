package users

import (
	"time"
)

type UserDomain struct {
	ID           uint
	Username     string
	Email        string
	Password     string
	HashPassword string
	Fullname     string
	Dob          time.Time
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	Signup(domain UserDomain) (UserDomain, error)
	Signin(username string, password string) (UserDomain, error)
}

type Repository interface {
	Signup(domain UserDomain) (UserDomain, error)
	Signin(username string) (UserDomain, error)
}
