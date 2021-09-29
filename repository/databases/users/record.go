package users

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
	"github.com/Rahmanwghazi/Monefy/repository/databases/income"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Fullname string
	Dob      time.Time
	Income   []income.Income
}

func (user *User) ToDomain() users.UserDomain {
	return users.UserDomain{
		Model:    gorm.Model{},
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Fullname: user.Fullname,
		Dob:      user.Dob,
	}
}

func FromDomain(domain users.UserDomain) User {
	return User{
		Model:    gorm.Model{},
		Username: domain.Username,
		Email:    domain.Email,
		Password: domain.Password,
		Fullname: domain.Fullname,
		Dob:      domain.Dob,
	}
}
