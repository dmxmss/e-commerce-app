package repository

import (
	"github.com/dmxmss/e-commerce-app/config"
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/golang-jwt/jwt/v5"

	"errors"
)

type AuthRepository interface {
	ValidateToken(string) (*entities.Claims, error)
	GenerateAndSignToken(entities.Claims) (string, error)
}

type authRepository struct {
	method jwt.SigningMethod
	secret string
}

func NewAuthRepository(conf *config.Auth) (AuthRepository, error) {
	var method jwt.SigningMethod

	switch conf.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256	
	default:
		return nil, e.AuthSignatureInvalid{}	
	}

	return &authRepository{
		method: method,
		secret: conf.JWTSecret,
	}, nil
}

func (r *authRepository) ValidateToken(rawToken string) (*entities.Claims, error) {
	token, err := jwt.ParseWithClaims(rawToken, &entities.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != r.method.Alg() {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(r.secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, e.AuthSignatureInvalid{}
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, e.TokenExpired{}
		} else {
			return nil, e.AuthFailed{}
		}
	}

	claims, ok := token.Claims.(*entities.Claims)

	if !ok || !token.Valid {
		return nil, e.TokenInvalid{}
	}

	return claims, nil
}


func (r *authRepository) signToken(token *jwt.Token) (string, error) {
	signedToken, err := token.SignedString([]byte(r.secret))
	if err != nil {
		return "", e.TokenSigningError{}
	}

	return signedToken, nil
}

func (r *authRepository) GenerateAndSignToken(claims entities.Claims) (string, error) {
	token := jwt.NewWithClaims(r.method, claims)	
	signedToken, err := r.signToken(token)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
