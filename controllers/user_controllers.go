package controllers

import (
	"net/http"

	"github.com/Rahmanwghazi/Monefy/config"
	"github.com/Rahmanwghazi/Monefy/middleware"
	"github.com/Rahmanwghazi/Monefy/models"
	"github.com/labstack/echo/v4"
)

func SignUpController(echoContext echo.Context) error {
	user := models.User{}
	echoContext.Bind(&user)
	err := config.InitDB().Save(&user).Error
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"message": "user created",
		"user":    user,
	})
}

func SignInController(echoContext echo.Context) error {
	user := models.User{}
	echoContext.Bind(&user)
	err := config.InitDB().Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.Username, user.Email)
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		})
	}
	userResponse := models.UserResponse{user.Username, user.Email, token}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"user":    userResponse,
	})
}
