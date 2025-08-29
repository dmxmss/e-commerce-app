package repository

import (
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	e "github.com/dmxmss/e-commerce-app/error"
	"gorm.io/gorm"
	
	"errors"
	"fmt"
)

type ProductRepository interface {
	CreateProduct(entities.Product) (*entities.Product, error)
	GetProductsBy(dto.GetProductsBy) ([]entities.Product, error)
	DeleteProduct(entities.Product) error
	GetCategoryByName(string) (*entities.Category, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) CreateProduct(product entities.Product) (*entities.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return nil, e.DbTransactionFailed{Err: errors.New("failed to create a product")}
	}

	return &product, nil
}

func (r *productRepository) GetProductsBy(request dto.GetProductsBy) ([]entities.Product, error) {
	var products []entities.Product
	query := r.db.Model(&products)

	if request.ID == nil && request.Name == "" && request.Vendor == nil && request.Category == "" {
		return nil, nil
	}

	if request.ID != nil {
		query = query.Where("id = ?", request.ID)
	}

	if request.Name != "" {
		query = query.Where("name = ?", request.Name)
	}

	if request.Vendor != nil {
		query = query.Where("vendor = ?", request.Vendor)
	}

	if request.Category != "" {
		category, err := r.GetCategoryByName(request.Category)
		if err != nil {
			return nil, err
		}

		query = query.Where("category = ?", category.ID)
	}

	if err := query.Find(&products).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.DbRecordNotFound{Err: "no products found with this condition"}
		} else {
			return nil, e.DbTransactionFailed{Err: err}
		}
	}

	return products, nil
}

func (r *productRepository) DeleteProduct(product entities.Product) error {
	if err := r.db.Delete(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.DbRecordNotFound{Err: fmt.Sprintf("product with id %d not found", product.ID)}
		} else {
			return e.DbTransactionFailed{Err: err}
		}
	}

	return nil
}

func (r *productRepository) GetCategoryByName(name string) (*entities.Category, error) {
	if name == "" {
		return nil, e.InvalidInputError{Err: "category name is empty"}
	}

	var category entities.Category

	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.DbRecordNotFound{Err: fmt.Sprintf("category with name '%s' not found", name)}
		} else {
			return nil, e.DbTransactionFailed{Err: err}
		}
	}

	return &category, nil
}
