package repository

import (
	"github.com/dmxmss/e-commerce-app/config"
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/golang-jwt/jwt/v5"
)

type AuthRepository interface {
	GenerateAndSignToken(entities.Claims) (*entities.AuthToken, error)
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

func (r *authRepository) signToken(token *jwt.Token) (string, error) {
	signedToken, err := token.SignedString([]byte(r.secret))
	if err != nil {
		return "", e.TokenSigningError{}
	}

	return signedToken, nil
}

func (r *authRepository) GenerateAndSignToken(claims entities.Claims) (*entities.AuthToken, error) {
	rawToken := jwt.NewWithClaims(r.method, claims)	
	signedToken, err := r.signToken(rawToken)
	if err != nil {
		return nil, err
	}
	
	token := entities.AuthToken{
		Value: signedToken,
		Claims: claims,
	}

	return &token, nil
}
