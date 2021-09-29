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

func GetUser(echoContext echo.Context) *JwtCustomClaims {
	user := echoContext.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

/* func ExtractJWT(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2 {
		return bearerToken[1]
	}
	return ""
}

func CheckToken(r *http.Request) (*jwt.Token, error) {
	tokenStr := ExtractJWT(r)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return " ", nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMeta(r *http.Request) (*JwtCustomClaims, error) {
	token, err := CheckToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, ok := claims["user_id"].(uint)
		if !ok {
			return nil, err
		}
		return &JwtCustomClaims{
			ID: userId,
		}, nil
	}
	return nil, err
}
*/
