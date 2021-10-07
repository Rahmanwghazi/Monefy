package users

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/users"
	"github.com/Rahmanwghazi/Monefy/repository/databases/expenses"
	"github.com/Rahmanwghazi/Monefy/repository/databases/income"
	"github.com/Rahmanwghazi/Monefy/repository/databases/investplans"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	HashPassword string
	Fullname     string
	Dob          time.Time
	Income       []income.Income
	Expense      []expenses.Expense
	InvestPlan   []investplans.InvestPlan
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (user *User) ToDomain() users.UserDomain {
	return users.UserDomain{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		HashPassword: user.HashPassword,
		Fullname:     user.Fullname,
		Dob:          user.Dob,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func FromDomain(domain users.UserDomain) *User {
	return &User{
		ID:           domain.ID,
		Username:     domain.Username,
		Email:        domain.Email,
		HashPassword: domain.HashPassword,
		Fullname:     domain.Fullname,
		Dob:          domain.Dob,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
