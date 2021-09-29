package users

import (
	"time"

	"gorm.io/gorm"
)

type UserDomain struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Fullname string
	Dob      time.Time
	Token    string
}

type Usecase interface {
	Signup(domain *UserDomain) (UserDomain, error)
	Signin(username string, password string) (UserDomain, error)
}

type Repository interface {
	Signup(domain *UserDomain) (UserDomain, error)
	Signin(username string, password string) (UserDomain, error)
}
