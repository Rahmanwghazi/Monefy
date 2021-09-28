package users

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
	"github.com/Rahmanwghazi/Monefy/drivers/databases/income"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	FullName string
	DoB      time.Time
	Income   []income.Income
}

func (user *User) ToDomain() users.UserDomain {
	return users.UserDomain{
		Model:    gorm.Model{},
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		FullName: user.FullName,
		DoB:      user.DoB,
	}
}

func FromDomain(domain users.UserDomain) User {
	return User{
		Model:    gorm.Model{},
		Username: domain.Username,
		Email:    domain.Email,
		Password: domain.Password,
		FullName: domain.FullName,
		DoB:      domain.DoB,
	}
}
