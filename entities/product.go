package entities

import (
	"time"
)

type Product struct {
	ID int `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Name string `gorm:"not null"`
	Description string
	Vendor string
	Price int 
	Tags string
}
