package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("a996dab369495e6b7d3a8bd19497a48a")
)

type CustomClaims struct {
	User *User
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(token string) (*CustomClaims, error) {
	tokenType, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(user *User) (string, error) {
	exp := time.Now().Add(time.Minute * time.Duration(10)).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    "shippy.service.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}
