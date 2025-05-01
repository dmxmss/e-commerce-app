package service

import (
	"github.com/dmxmss/e-commerce-app/config"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"github.com/golang-jwt/jwt/v5"

	"time"
)

type AuthService interface {
	GenerateToken(string, int) (string, error)
	GenerateTokens(string) (string, string, error)
}

type authService struct {
	authRepo repository.AuthRepository
	accessExpiration int
	refreshExpiration int
}

func NewAuthService(conf *config.Config) (AuthService, error) {
	authRepo, err := repository.NewAuthRepository(conf.Auth)	
	if err != nil {
		return nil, err
	}

	return &authService{
		authRepo: authRepo,
		accessExpiration: conf.Auth.Access.Expiration,
		refreshExpiration: conf.Auth.Refresh.Expiration,
	}, nil
}

func (s *authService) GenerateToken(userId string, expiration int) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject: userId,	
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiration)*time.Second)),
	}
	token, err := s.authRepo.GenerateAndSignToken(claims)
	
	return token, err
}

func (s *authService) GenerateTokens(userId string) (string, string, error) {
	access, err := s.GenerateToken(userId, s.accessExpiration)
	if err != nil {
		return "", "", err
	}

	refresh, err := s.GenerateToken(userId, s.refreshExpiration)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}
