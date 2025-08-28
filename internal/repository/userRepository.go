package repository

import (
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"gorm.io/gorm"

	"errors"
)

type UserRepository interface {
	CreateUser(user entities.User) (*entities.User, error)
	GetUserBy(dto.GetUserBy) (*entities.User, error)
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
			return nil, e.DbTransactionFailed{Err: err}
		}
	}

	return &user, nil
}

func (u *userRepository) GetUserBy(request dto.GetUserBy) (*entities.User, error) { // only one field in request must be non nil
	var user entities.User	
	query := u.db.Model(&user)

	if request.Name == "" && request.Email == "" && request.ID == nil {
		return nil, nil
	}

	if request.Name != "" {
		query = query.Where("name = ?", request.Name)		
	}

	if request.Email != "" {
		query = query.Where("email = ?", request.Email)		
	}

	if request.ID != nil {
		query = query.Where("id = ?", request.ID)
	}

	if err := query.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.UserNotFound{Name: request.Name}
		} else {
			return nil, e.DbTransactionFailed{Err: err}
		}
	}

	return &user, nil
}
