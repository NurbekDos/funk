package services

import (
	"github.com/NurbekDos/funk/internal/cfg"
	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	UserId uint
	Email  string
	jwt.StandardClaims
}

func GenerateToken(claims TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.GetConfig().TokenKey))
}

func VerifyToken(token string) *TokenClaims {
	parsedToken, _ := jwt.ParseWithClaims(token, &TokenClaims{}, func(tokenX *jwt.Token) (interface{}, error) {
		return []byte(cfg.GetConfig().TokenKey), nil
	})

	return parsedToken.Claims.(*TokenClaims)
}
