package requests

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
)

type User struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Fullname string    `json:"fullname"`
	Dob      time.Time `json:"dob"`
}

func (User *User) ToDomain() users.UserDomain {
	return users.UserDomain{
		Username: User.Username,
		Email:    User.Email,
		Password: User.Password,
		Fullname: User.Fullname,
		Dob:      User.Dob,
	}
}
