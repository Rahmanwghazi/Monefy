package signin

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
)

type UserSignin struct {
	//gorm.Model
	Username string    `json:"username" gorm:"unique"`
	Email    string    `json:"email"  gorm:"unique"`
	Password string    `json:"password"`
	FullName string    `json:"fullname"`
	DoB      time.Time `json:"dob"`
	//Token    string    `json:"token"`
}

func FromDomain(domain users.Domain) UserSignin {
	return UserSignin{
		//Model:    gorm.Model{},
		Username: domain.Username,
		Email:    domain.Email,
		FullName: domain.FullName,
		DoB:      domain.DoB,
		//Token:    domain.Token,
	}
}
