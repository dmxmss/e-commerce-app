package repository

import (
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"gorm.io/gorm"
	
	"errors"
)

type ProductRepository interface {
	CreateProduct(entities.Product) (*entities.Product, error)
	DeleteProduct(entities.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) CreateProduct(product entities.Product) (*entities.Product, error) {
	if err := p.db.Create(&product).Error; err != nil {
		return nil, e.DbTransactionFailed{Err: errors.New("failed to create a product")}
	}

	return &product, nil
}

func (p *productRepository) DeleteProduct(product entities.Product) error {
	if err := p.db.Delete(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.ProductNotFound{ID: product.ID}
		} else {
			return e.DbTransactionFailed{Err: err}
		}
	}

	return nil
}
