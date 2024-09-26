package services

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	UserId uint
	Email  string
	jwt.StandardClaims
}

var tokenKey = "random_string" // TODO to config

func GenerateToken(claims TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(tokenKey))
}

func VerifyToken(token string) *TokenClaims {
	parsedToken, _ := jwt.ParseWithClaims(token, &TokenClaims{}, func(tokenX *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})

	return parsedToken.Claims.(*TokenClaims)
}
