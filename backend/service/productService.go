package service

import (
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/internal/repository"
)

type ProductService interface {
	CreateProduct(string, string, int, int, int, string) (*entities.Product, error)
	GetProducts(dto.GetProductParams) ([]entities.Product, error)
	GetProduct(int) (*entities.Product, error)
	UpdateProduct(int, dto.UpdateProductRequest) (*entities.Product, error)
	DeleteProduct(int) error
}

type productServiceRepo struct { // repositories product service needs
	product repository.ProductRepository
}

type productService struct {
	repo productServiceRepo
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
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

func (s *productService) GetProducts(params dto.GetProductParams) ([]entities.Product, error) {
	return s.repo.product.GetProducts(params)
}

func (s *productService) GetProduct(id int) (*entities.Product, error) {
	return s.repo.product.GetProduct(id)
}

func (s *productService) UpdateProduct(id int, request dto.UpdateProductRequest) (*entities.Product, error) {
	return s.repo.product.UpdateProduct(id, request)
}

func (s *productService) DeleteProduct(id int) error {
	return s.repo.product.DeleteProduct(id)
}
