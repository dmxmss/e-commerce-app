package service

import (
	"github.com/dmxmss/e-commerce-app/config"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/dmxmss/e-commerce-app/internal/dto"

	"time"
	"strconv"
)

type AuthService interface {
	SignUp(string, string, string) (*entities.User, *entities.AuthToken, *entities.AuthToken, error)
	Login(dto.LoginRequest) (*entities.User, *entities.AuthToken, *entities.AuthToken, error)
	GenerateTokens(bool, string) (*entities.AuthToken, *entities.AuthToken, error)
}

type authServiceRepo struct { // repositories auth service needs
	auth repository.AuthRepository
	user repository.UserRepository
}

type authService struct {
	repo authServiceRepo
	conf *config.Auth
}

func NewAuthService(
	conf *config.Config, 
	authRepo repository.AuthRepository, 
	userRepo repository.UserRepository,
) (AuthService, error) {

	return &authService{
		repo: authServiceRepo{
			auth: authRepo,
			user: userRepo,
		},
		conf: conf.Auth,
	}, nil
}

func (s *authService) SignUp(name, email, password string) (*entities.User, *entities.AuthToken, *entities.AuthToken, error) {
	request := entities.User{
		Name: name,
		Email: email, 
		Password: password,
	}	

	user, err := s.repo.user.CreateUser(request)
	if err != nil {
		return nil, nil, nil, err
	}

	accessToken, refreshToken, err := s.GenerateTokens(false, strconv.Itoa(user.ID))
	if err != nil {
		return nil, nil, nil, err
	}

	return user, accessToken, refreshToken, nil
}

func (s *authService) Login(login dto.LoginRequest) (*entities.User, *entities.AuthToken, *entities.AuthToken, error) {
	params := dto.GetUsersParams{
		Name: login.Name,
		Email: login.Email,
	}

	users, total, err := s.repo.user.GetUsers(params)
	if err != nil {
		return nil, nil, nil, err
	}

	if total != 1 {
		return nil, nil, nil, e.InternalServerError{Err: "auth service: login: total != 1"}
	}

	user := users[0]

	if user.Password != login.Password {
		return nil, nil, nil, e.InvalidCredentials{}
	}

	accessToken, refreshToken, err := s.GenerateTokens(false, strconv.Itoa(user.ID))
	if err != nil {
		return nil, nil, nil, err
	}

	return &user, accessToken, refreshToken, nil
}

func (s *authService) GenerateTokens(isAdmin bool, subject string) (*entities.AuthToken, *entities.AuthToken, error) {
	accessTokenClaims := entities.Claims {
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.conf.Access.Expiration)*time.Second)),
			Subject: subject,
		},
	}

	accessToken, err := s.repo.auth.GenerateAndSignToken(accessTokenClaims)
	if err != nil {
		return nil, nil, err
	}

	refreshTokenClaims := entities.Claims {
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.conf.Refresh.Expiration)*time.Second)),
			Subject: subject,
		},
	}

	refreshToken, err := s.repo.auth.GenerateAndSignToken(refreshTokenClaims)
	if err != nil {
		return nil, nil, err
	}
	
	return accessToken, refreshToken, nil
}
