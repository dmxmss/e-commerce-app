package entities

import (
	"time"
)

type Product struct {
	ID int `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Name string `gorm:"not null"`
	Description string
	Vendor int
	User User `gorm:"foreignKey:Vendor"`
	Remaining int
	Price int 
	CategoryID int
	Category Category
	Images []ProductImage `gorm:"foreignKey:ProductID"`
}
