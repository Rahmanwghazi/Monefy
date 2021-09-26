package requests

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
	//"gorm.io/gorm"
)

type UserSignup struct {
	//gorm.Model
	Username string    `json:"username" gorm:"unique"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password"`
	FullName string    `json:"fullname"`
	DoB      time.Time `json:"dob"`
}

func (UserSignup *UserSignup) ToDomain() users.Domain {
	return users.Domain{
		Username: UserSignup.Username,
		Email:    UserSignup.Email,
		Password: UserSignup.Password,
		FullName: UserSignup.FullName,
		DoB:      UserSignup.DoB,
	}
}
