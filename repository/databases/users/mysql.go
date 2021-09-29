package users

import (
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

func (rep *mysqlUserRepository) Signup(domain *users.UserDomain) (users.UserDomain, error) {
	user := User{}
	user.Username = domain.Username
	user.Email = domain.Email
	user.Password = domain.Password
	user.Fullname = domain.Fullname
	user.Dob = domain.Dob

	result := rep.Connection.Create(&user)

	if result.Error != nil {
		return users.UserDomain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (repository *mysqlUserRepository) Signin(username string, password string) (users.UserDomain, error) {
	var user User
	err := repository.Connection.First(&user, "username = ? AND password = ?", username, password).Error

	if err != nil {
		return users.UserDomain{}, err
	}

	return user.ToDomain(), nil
}
