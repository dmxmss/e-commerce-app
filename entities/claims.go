package entities

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int
	jwt.RegisteredClaims
}
