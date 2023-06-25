package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// token oluşturma

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})

	return claims.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, error) {
	// cooki yi pars etme
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, err
}
