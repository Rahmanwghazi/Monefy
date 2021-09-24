package middleware

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/constants"
	"github.com/golang-jwt/jwt"
)

func CreateToken(email string, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["username"] = username
	claims["expire"] = time.Now().Add(time.Hour * 3).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
