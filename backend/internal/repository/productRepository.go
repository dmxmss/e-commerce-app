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
	GetProducts([]int) ([]entities.Product, error)
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

	if request.Names == nil && request.Vendor == nil && request.Categories == nil {
		return nil, nil
	}

	if request.Names != nil {
		query = query.Where("name IN ?", request.Names)
	}

	if request.Vendor != nil {
		query = query.Where("vendor = ?", request.Vendor)
	}

	if request.Categories != nil {
		query = query.Where("category = ?", request.Categories)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, e.DbTransactionFailed{Err: err}
	}

	if len(products) == 0 {
		return nil, e.DbRecordNotFound{Err: "no records found with these conditions"}
	}

	return products, nil
}

func (r *productRepository) GetProducts(ids []int) ([]entities.Product, error) {
	var products []entities.Product

	if err := r.db.Where("id IN ?", ids).Find(&products).Error; err != nil {
		return nil, e.DbTransactionFailed{Err: err}
	}

	if len(products) != len(ids) {
		return nil, e.DbRecordNotFound{Err: "not all records with these ids are found"}
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
