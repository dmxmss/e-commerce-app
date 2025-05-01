package repository

import (
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"gorm.io/gorm"

	"errors"
)

type UserRepository interface {
	CreateUser(dto.CreateUserRequest) (*entities.User, error)
	GetUserBy(dto.GetUserRequest) (*entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(createUser dto.CreateUserRequest) (*entities.User, error) {
	user := entities.User{
		Name: createUser.Name, 
		Email: createUser.Email,
		Password: createUser.Password,
	}

	if err := u.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, e.UserAlreadyExists{Name: createUser.Name}
		} else {
			return nil, e.DbTransactionFailed{Err: err}
		}
	}

	return &user, nil
}

func (u *userRepository) GetUserBy(getUser dto.GetUserRequest) (*entities.User, error) {
	var user entities.User	
	query := u.db.Model(&user)

	if getUser.Name == "" && getUser.Email == "" && getUser.ID == nil {
		return nil, nil
	}

	if getUser.ID != nil {
		query = query.Where("id = ?", getUser.ID)
	}

	if getUser.Name != "" {
		query = query.Where("id = ?", getUser.Name)		
	}

	if getUser.Email != "" {
		query = query.Where("email = ?", getUser.Email)		
	}

	if err := query.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.UserNotFound{Name: getUser.Name}
		} else {
			return nil, e.DbTransactionFailed{Err: err}
		}
	}

	return &user, nil
}
