package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id ,secret string) (string, error) {
	claims := &JwtClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72*time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err  := token.SignedString([]byte(secret))
	if err != nil{
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString, secret string) (*JwtClaims, error){
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims , ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid{
		return nil, errors.New("invalid token")
	}
	return claims, nil
}