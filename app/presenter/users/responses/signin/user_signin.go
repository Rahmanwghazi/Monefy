package signin

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
)

type UserSignin struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullname"`
	DoB       time.Time `json:"dob"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.UserDomain) UserSignin {
	return UserSignin{
		ID:        domain.ID,
		Username:  domain.Username,
		Email:     domain.Email,
		FullName:  domain.Fullname,
		DoB:       domain.Dob,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
