package entities

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
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

func (p Product) ToResponse() dto.Product {
	var images []string
	for _, image := range p.Images {
		images = append(images, image.URL)
	}

	return dto.Product{
		ID: p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Name: p.Name,
		Description: p.Description,
		Vendor: p.Vendor,
		Remaining: p.Remaining,
		Price: p.Price,
		Category: p.CategoryID,
		Images: images,
	}
}
