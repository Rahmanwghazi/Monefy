package middlewares

import (
	"net/http"
	"time"

	controllers "github.com/Rahmanwghazi/Monefy/app/presenter"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID uint `json:"id"`
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
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(err error, echoContext echo.Context) error {
			return controllers.NewErrorResponse(echoContext, http.StatusForbidden, err)
		}),
	}
}

func (configJWT *ConfigJWT) GenerateTokenJWT(userId uint) (string, error) {
	claims := JwtCustomClaims{
		userId,
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

func ExtractClaims(echoContext echo.Context) (*JwtCustomClaims, error) {
	user := echoContext.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims, nil
}
