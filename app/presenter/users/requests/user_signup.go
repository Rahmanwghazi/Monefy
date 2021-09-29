package requests

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
)

type UserSignup struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Fullname string    `json:"fullname"`
	Dob      time.Time `json:"dob"`
}

func (UserSignup *UserSignup) ToDomain() users.UserDomain {
	return users.UserDomain{
		Username: UserSignup.Username,
		Email:    UserSignup.Email,
		Password: UserSignup.Password,
		Fullname: UserSignup.Fullname,
		Dob:      UserSignup.Dob,
	}
}
