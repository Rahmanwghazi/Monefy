package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string     `json:"username" form:"username" gorm:"unique"`
	Email    string     `json:"email" form:"email" gorm:"unique"`
	Password string     `json:"password" form:"password"`
	FullName string     `json:"fullname" form:"fullname"`
	DoB      *time.Time `json:"dob" form:"dob"`
}

type UserResponse struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}
