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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token"`
}

func FromDomain(domain users.UserDomain) UserSignin {
	return UserSignin{
		ID:        domain.ID,
		Username:  domain.Username,
		Email:     domain.Email,
		FullName:  domain.Fullname,
		DoB:       domain.Dob,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		Token:     domain.Token,
	}
}
