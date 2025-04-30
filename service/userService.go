package service

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(dto.CreateUserRequest) (*dto.GetUserResponse, error)
	LogIn(dto.LoginUserRequest) (*dto.GetUserResponse, error)
	GetUserInfo(int) (*dto.GetUserResponse, error)
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
	}

	return &response, nil
}

func (s *userService) LogIn(login dto.LoginUserRequest) (*dto.GetUserResponse, error) {
	user, err := s.userRepo.GetUserBy(dto.GetUserRequest{Email: login.Email})
	if err != nil {
		return nil, err
	}

	response := dto.GetUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return &response, nil
}

func (s *userService) GetUserInfo(userId int) (*dto.GetUserResponse, error) {
	user, err := s.userRepo.GetUserBy(dto.GetUserRequest{ID: &userId})
	if err != nil {
		return nil, err
	}

	response := dto.GetUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return &response, nil
}
