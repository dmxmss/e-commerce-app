package repository

import (
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"errors"
	"fmt"
)

type ProductRepository interface {
	CreateProduct(entities.Product) (*entities.Product, error)
	GetProducts(dto.GetProductParams) ([]entities.Product, error)
	GetProduct(int) (*entities.Product, error)
	UpdateProduct(int, dto.UpdateProductRequest) (*entities.Product, error)
	DeleteProduct(int) error
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

func (r *productRepository) GetProduct(id int) (*entities.Product, error) {
	var product entities.Product

	if err := r.db.Where("id = ?", id).Take(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.DbRecordNotFound{Err: fmt.Sprintf("not found product with id %d", id)}
		} else {
			return nil, e.DbTransactionFailed{Err: err}
		}
	}

	return &product, nil
}

func (r *productRepository) GetProducts(params dto.GetProductParams) ([]entities.Product, error) {
	var products []entities.Product

	q := r.db.Model(&entities.Product{})

	if params.Target == "vendorId" && params.ID != 0 {
		q = q.Where("vendor = ?", params.ID)
	} else if params.IDs != nil {
		q = q.Where("id IN ?", params.IDs)
	} 

	// TODO: add filtering by price, remaining, category, created and updated date
	// TODO: add allowed sorting fields 

	if params.SortField != "" && params.SortOrder != "" {
		q = q.Order(params.SortField + " " + params.SortOrder)
	}

	if params.Page != 0 && params.PerPage != 0 {
		q = q.Limit(params.PerPage).Offset((params.Page - 1)*params.PerPage)
	}

	if err := q.Find(&products).Error; err != nil {
		return nil, e.DbTransactionFailed{Err: err}
	}

	return products, nil
}

func (r *productRepository) UpdateProduct(id int, request dto.UpdateProductRequest) (*entities.Product, error) {
	var product entities.Product

	if err := r.db.Model(&entities.Product{}).
								 Clauses(clause.Returning{}).
								 Where("id = ?", id).
								 Updates(&request).
								 Scan(&product).Error; err != nil {
		return nil, e.DbTransactionFailed{Err: err}
	}

	product.ID = id

	return &product, nil
}

func (r *productRepository) DeleteProduct(id int) error {
	product := entities.Product{
		ID: id,
	}

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
