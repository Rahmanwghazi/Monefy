package signup

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
)

type UserSignup struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Fullname string    `json:"fullname"`
	DoB      time.Time `json:"dob"`
}

func FromDomain(domain users.UserDomain) UserSignup {
	return UserSignup{
		Username: domain.Username,
		Email:    domain.Email,
		Fullname: domain.Fullname,
		DoB:      domain.Dob,
	}
}
