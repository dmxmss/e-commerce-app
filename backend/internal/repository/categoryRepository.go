package repository

import (
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	e "github.com/dmxmss/e-commerce-app/error"
	"gorm.io/gorm"

	"errors"
)

type CategoryRepository interface {
	GetCategory(int) (*entities.Category, error)
	GetCategories(dto.GetCategoriesParams) ([]entities.Category, int64, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetCategory(id int) (*entities.Category, error) {
	var category entities.Category

	if err := r.db.Find(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.DbRecordNotFound{Err: "category not found"}
		} else {
			return nil, e.DbTransactionFailed{Err: err.Error()}
		}
	}

	return &category, nil
}

func (r *categoryRepository) GetCategories(params dto.GetCategoriesParams) ([]entities.Category, int64, error) {
	var categories []entities.Category
	var total int64

	q := r.db.Model(&entities.Category{})

	if params.IDs != nil {
		q = q.Where("id IN ?", params.IDs)
	}

	handleSorting(q, params.SortField, params.SortOrder)

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, e.DbTransactionFailed{Err: err.Error()}
	}

	handlePagination(q, params.Page, params.PerPage)

	if err := q.Find(&categories).Error; err != nil {
		return nil, 0, e.DbTransactionFailed{Err: err.Error()}
	}

	return categories, total, nil
}
