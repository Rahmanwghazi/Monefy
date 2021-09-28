package middlewares

import (
	"net/http"
	"time"

	"github.com/Rahmanwghazi/Monefy/controllers"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJwt       string
	ExpiredDuration int
}

func (configJWT *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(configJWT.SecretJwt),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(err error, context echo.Context) error {
			return controllers.NewErrorResponse(context, http.StatusForbidden, err)
		}),
	}
}

func (configJWT *ConfigJWT) GenerateTokenJWT(id int) (string, error) {
	claims := JwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(configJWT.ExpiredDuration))).Unix(),
		},
	}

	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := createToken.SignedString([]byte(configJWT.SecretJwt))

	if err != nil {
		return "", err
	}
	return token, err
}
