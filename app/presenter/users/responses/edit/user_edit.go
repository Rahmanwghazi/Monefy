package edit

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
)

type UserEdit struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	DoB       time.Time `json:"dob"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.UserDomain) UserEdit {
	return UserEdit{
		ID:        domain.ID,
		Username:  domain.Username,
		Email:     domain.Email,
		Fullname:  domain.Fullname,
		DoB:       domain.Dob,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
