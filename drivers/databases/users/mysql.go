package users

import (
	"context"

	"github.com/Rahmanwghazi/Monefy/business/users"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Connection *gorm.DB
}

func NewMysqlUserRepository(connection *gorm.DB) users.Repository {
	return &mysqlUserRepository{
		Connection: connection,
	}
}

func (rep *mysqlUserRepository) Signup(context context.Context, domain users.Domain) (users.Domain, error) {
	user := User{}
	user.Username = domain.Username
	user.Email = domain.Email
	user.Password = domain.Password
	user.FullName = domain.FullName
	user.DoB = domain.DoB

	result := rep.Connection.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (repository *mysqlUserRepository) Signin(context context.Context, username string, password string) (users.Domain, error) {
	var user User
	err := repository.Connection.First(&user, "username = ? AND password = ?", username, password).Error

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}
