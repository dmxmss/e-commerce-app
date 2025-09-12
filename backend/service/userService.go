package service

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/repository"
)

type UserService interface {
	GetUsers(dto.GetUsersParams) ([]entities.User, int64, error)
	GetUser(int) (*entities.User, error)
}

type userServiceRepo struct { // repositories user service needs
	user repository.UserRepository
}

type userService struct {
	repo userServiceRepo
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		repo: userServiceRepo{
			user: userRepo,
		},
	}
}

func (s *userService) GetUsers(params dto.GetUsersParams) ([]entities.User, int64, error) {
	return s.repo.user.GetUsers(params)
}

func (s *userService) GetUser(id int) (*entities.User, error) {
	return s.repo.user.GetUser(id)
}
