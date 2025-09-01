package service

import (
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(string, string, int, int, int, string) (*entities.Product, error)
	GetUserProducts(int) ([]entities.Product, error)
	DeleteProduct(int) error
}

type productServiceRepo struct { // repositories product service needs
	product repository.ProductRepository
}

type productService struct {
	repo productServiceRepo
}

func NewProductService(db *gorm.DB) ProductService {
	productRepo := repository.NewProductRepository(db)

	return &productService{
		repo: productServiceRepo{
			product: productRepo,
		},
	}
}

func (s *productService) CreateProduct(name, description string, vendor, remaining, price int, categoryName string) (*entities.Product, error) {
	category, err := s.repo.product.GetCategoryByName(categoryName)
	if err != nil {
		return nil, err
	}

	product := entities.Product{
		Name: name,
		Description: description,
		Vendor: vendor,
		Remaining: remaining,
		Price: price,
		Category: *category,
	}

	response, err := s.repo.product.CreateProduct(product)

	return response, err
}

func (s *productService) DeleteProduct(id int) error {
	product := entities.Product{
		ID: id,
	}

	return s.repo.product.DeleteProduct(product)
}

func (s *productService) GetUserProducts(id int) ([]entities.Product, error) {
	products, err := s.repo.product.GetProductsBy(dto.GetProductsBy{Vendor: &id})
	if err != nil {
		return nil, err
	}

	return products, nil
}
