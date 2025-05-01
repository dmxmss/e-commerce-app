package service

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"gorm.io/gorm"

	"strconv"
)

type UserService interface {
	CreateUser(dto.CreateUserRequest) (*dto.GetUserResponse, error)
	LogIn(dto.LoginUserRequest) (*dto.GetUserResponse, error)
	GetUserInfo(string) (*dto.GetUserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	userRepo := repository.NewUserRepository(db)

	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(createUser dto.CreateUserRequest) (*dto.GetUserResponse, error) {
	user, err := s.userRepo.CreateUser(createUser)
	if err != nil {
		return nil, err
	}

	response := dto.GetUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Admin: user.Admin,
	}

	return &response, nil
}

func (s *userService) LogIn(login dto.LoginUserRequest) (*dto.GetUserResponse, error) {
	user, err := s.userRepo.GetUserBy(dto.GetUserRequest{Email: login.Email})
	if err != nil {
		return nil, err
	}

	if login.Password != user.Password {
		return nil, e.InvalidCredentials{}
	}

	response := dto.GetUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Admin: user.Admin,
	}

	return &response, nil
}

func (s *userService) GetUserInfo(userId string) (*dto.GetUserResponse, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, e.InvalidUserId{ID: userId}
	}

	user, err := s.userRepo.GetUserBy(dto.GetUserRequest{ID: &id})
	if err != nil {
		return nil, err
	}

	response := dto.GetUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Admin: user.Admin,
	}

	return &response, nil
}
