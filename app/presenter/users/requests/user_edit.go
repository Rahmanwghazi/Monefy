package requests

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
)

type UserEdit struct {
	ID       uint      `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Fullname string    `json:"fullname"`
	Dob      time.Time `json:"dob"`
}

func (UserEdit *UserEdit) ToDomain() users.UserDomain {
	return users.UserDomain{
		ID:       UserEdit.ID,
		Username: UserEdit.Username,
		Email:    UserEdit.Email,
		Password: UserEdit.Password,
		Fullname: UserEdit.Fullname,
		Dob:      UserEdit.Dob,
	}
}
