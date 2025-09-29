package entities

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
)

type User struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Admin bool `gorm:"default:false"`
}

func (u User) ToResponse() dto.User {
	return dto.User{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
	}
}
