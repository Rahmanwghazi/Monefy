package signin

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
)

type UserSignin struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	FullName string    `json:"fullname"`
	DoB      time.Time `json:"dob"`
	Token    string    `json:"token"`
}

func FromDomain(domain users.UserDomain) UserSignin {
	return UserSignin{
		Username: domain.Username,
		Email:    domain.Email,
		FullName: domain.FullName,
		DoB:      domain.DoB,
		Token:    domain.Token,
	}
}
