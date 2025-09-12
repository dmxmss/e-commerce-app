package repository

import (
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"gorm.io/gorm"

	"errors"
	"fmt"
)

type UserRepository interface {
	CreateUser(user entities.User) (*entities.User, error)
	GetUser(int) (*entities.User, error)
	GetUsers(dto.GetUsersParams) ([]entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(user entities.User) (*entities.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, e.UserAlreadyExists{}
		} else {
			return nil, e.DbTransactionFailed{Err: err.Error()}
		}
	}

	return &user, nil
}

func (r *userRepository) GetUser(id int) (*entities.User, error) {
	var user entities.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.DbRecordNotFound{Err: fmt.Sprintf("user with id %d not found", id)}
		} else {
			return nil, e.DbTransactionFailed{Err: err.Error()}
		}
	}

	return &user, nil
}

func (r *userRepository) GetUsers(params dto.GetUsersParams) ([]entities.User, error) {
	var users []entities.User	
	q:= r.db.Model(&entities.User{})

	if params.IDs != nil {
		q = q.Where("id IN ?", params.IDs)
	}

	if params.Name != "" {
		q = q.Where("name = ?", params.Name)
	}
	
	if params.Email != "" {
		q = q.Where("email = ?", params.Email)
	}

	if params.SortField != "" && params.SortOrder != "" {
		q = q.Order(params.SortField + " " + params.SortOrder)
	}

	if params.Page != 0 && params.PerPage != 0 {
		q = q.Limit(params.PerPage).Offset((params.Page - 1)*params.PerPage)
	}

	if err := q.Find(&users).Error; err != nil {
		return nil, e.DbTransactionFailed{Err: err.Error()}
	}

	return users, nil
}
