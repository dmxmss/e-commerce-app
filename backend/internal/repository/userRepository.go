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
	GetUsers(dto.GetUsersParams) ([]entities.User, int64, error)
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

func (r *userRepository) GetUsers(params dto.GetUsersParams) ([]entities.User, int64, error) {
	var users []entities.User	
	var total int64

	q := r.db.Model(&entities.User{})

	if params.IDs != nil {
		q = q.Where("id IN ?", params.IDs)
	}

	if params.Name != "" {
		q = q.Where("name = ?", params.Name)
	}
	
	if params.Email != "" {
		q = q.Where("email = ?", params.Email)
	}

	allowedFields := []string{"", "id", "name", "email"}

	if err := handleSorting(q, params.SortField, params.SortOrder, allowedFields); err != nil {
		return nil, 0, err
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, e.DbTransactionFailed{Err: err.Error()}
	}

	handlePagination(q, params.Page, params.PerPage)

	if err := q.Find(&users).Error; err != nil {
		return nil, 0, e.DbTransactionFailed{Err: err.Error()}
	}

	return users, total, nil
}
