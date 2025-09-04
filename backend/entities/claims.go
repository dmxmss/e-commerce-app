package entities

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Admin bool `json:"admin"`
	jwt.RegisteredClaims
} 
