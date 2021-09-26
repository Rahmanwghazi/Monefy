package signup

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
)

type UserSignup struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	FullName string    `json:"fullname"`
	DoB      time.Time `json:"dob"`
}

func FromDomain(domain users.Domain) UserSignup {
	return UserSignup{
		Username: domain.Username,
		Email:    domain.Email,
		FullName: domain.FullName,
		DoB:      domain.DoB,
	}
}
