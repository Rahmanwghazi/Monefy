package middleware

import (
	"github.com/Rahmanwghazi/Monefy/config"
	"github.com/Rahmanwghazi/Monefy/models"
	"github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, echoContext echo.Context) (bool, error) {
	var user models.User
	err := config.InitDB().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
