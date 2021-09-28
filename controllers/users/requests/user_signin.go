package requests

import "github.com/Rahmanwghazi/Monefy/business/users"

type UserSignin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (userSignin *UserSignin) ToDomain() users.UserDomain {
	return users.UserDomain{
		Email:    userSignin.Username,
		Password: userSignin.Password,
	}
}
