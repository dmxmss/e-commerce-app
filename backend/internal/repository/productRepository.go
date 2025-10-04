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
	GetProducts(dto.GetProductParams) ([]entities.Product, int64, error)
	GetProduct(int) (*entities.Product, error)
	UpdateProduct(int, dto.UpdateProductRequest) (*entities.Product, error)
	DeleteProduct(int) error
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
		return nil, e.DbTransactionFailed{Err: "failed to create a product"}
	}

	return &product, nil
}

func (r *productRepository) GetProduct(id int) (*entities.Product, error) {
	var product entities.Product

	if err := r.db.Preload("Images").Where("id = ?", id).Take(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.DbRecordNotFound{Err: fmt.Sprintf("not found product with id %d", id)}
		} else {
			return nil, e.DbTransactionFailed{Err: err.Error()}
		}
	}

	return &product, nil
}

func (r *productRepository) GetProducts(params dto.GetProductParams) ([]entities.Product, int64, error) {
	var products []entities.Product
	var total int64

	q := r.db.Preload("Images").Model(&entities.Product{})

	if params.Target == "vendorId" && params.ID != 0 {
		q = q.Where("vendor = ?", params.ID)
	} else if params.IDs != nil {
		q = q.Where("id IN ?", params.IDs)
	} 

	// TODO: add filtering by price, remaining, category, created and updated date
	if params.PriceMax != 0 {
		q = q.Where("price >= ?", params.PriceMax)
	}
	if params.PriceMin != 0 {
		q = q.Where("price <= ?", params.PriceMin)
	}

	if params.CreatedAfter.Time != nil {
		q = q.Where("created_at >= ?", params.CreatedAfter.Time)
	}
	if params.CreatedBefore.Time != nil {
		q = q.Where("created_at <= ?", params.CreatedBefore.Time)
	}
	if params.UpdatedAfter.Time != nil {
		q = q.Where("updated_at >= ?", params.UpdatedAfter.Time)
	}
	if params.UpdatedBefore.Time != nil {
		q = q.Where("updated_at <= ?", params.UpdatedBefore.Time)
	}

	if params.IsRemaining {
		q = q.Where("remaining > 0")
	}

	allowedFields := []string{"", "id", "name", "created_at", "updated_at", "remaining", "price"}

	if err := handleSorting(q, params.SortField, params.SortOrder, allowedFields); err != nil {
		return nil, 0, err
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, e.DbTransactionFailed{Err: err.Error()}
	}

	handlePagination(q, params.Page, params.PerPage)

	if err := q.Find(&products).Error; err != nil {
		return nil, 0, e.DbTransactionFailed{Err: err.Error()}
	}

	return products, total, nil
}

func (r *productRepository) UpdateProduct(id int, request dto.UpdateProductRequest) (*entities.Product, error) {
	var product entities.Product

	if err := r.db.Model(&entities.Product{}).
								 Clauses(clause.Returning{}).
								 Where("id = ?", id).
								 Updates(&request).
								 Scan(&product).Error; err != nil {
		return nil, e.DbTransactionFailed{Err: err.Error()}
	}

	product.ID = id

	return &product, nil
}

func (r *productRepository) DeleteProduct(id int) error {
	product := entities.Product{
		ID: id,
	}

	res := r.db.Delete(&product)
	err := res.Error
	if err != nil {
		return e.DbTransactionFailed{Err: err.Error()}
	}

	if res.RowsAffected == 0 {
		return e.DbRecordNotFound{Err: "product not found"}
	}

	return nil
}
