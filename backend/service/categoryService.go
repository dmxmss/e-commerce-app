package service

import (
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/internal/repository"
)

type CategoryService interface {
	GetCategory(int) (*entities.Category, error)
	GetCategories(dto.GetCategoriesParams) ([]entities.Category, int64, error)
}

type categoryServiceRepo struct { // repositories category service needs
	category repository.CategoryRepository
}

type categoryService struct {
	repo categoryServiceRepo
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: categoryServiceRepo{
			category: categoryRepo,
		},
	}
}

func (s *categoryService) GetCategory(id int) (*entities.Category, error) {
	return s.repo.category.GetCategory(id)
}

func (s *categoryService) GetCategories(params dto.GetCategoriesParams) ([]entities.Category, int64, error) {
	return s.repo.category.GetCategories(params)
}
