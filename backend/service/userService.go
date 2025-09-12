package service

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/repository"

	"strconv"
)

type UserService interface {
	GetUserInfo(string) (*entities.User, error)
	GetUsers(dto.GetUsersParams) ([]entities.User, error)
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

func (s *userService) GetUserInfo(userId string) (*entities.User, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, e.InvalidUserId{ID: userId}
	}

	user, err := s.repo.user.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUsers(params dto.GetUsersParams) ([]entities.User, error) {
	return s.repo.user.GetUsers(params)
}

func (s *userService) GetUser(id int) (*entities.User, error) {
	return s.repo.user.GetUser(id)
}
