package service

import (
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"gorm.io/gorm"

	"strings"
)

type ProductService interface {
	CreateProduct(dto.CreateProductRequest) (*entities.Product, error)
	DeleteProduct(dto.DeleteProductRequest) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(db *gorm.DB) ProductService {
	productRepo := repository.NewProductRepository(db)

	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateProduct(createProduct dto.CreateProductRequest) (*entities.Product, error) {
	concatTags := strings.Join(createProduct.Tags, ",")

	product := entities.Product{
		Name: createProduct.Name,
		Description: createProduct.Description,
		Vendor: createProduct.Vendor,
		Price: createProduct.Price,
		Tags: concatTags,
	}

	response, err := s.productRepo.CreateProduct(product)

	return response, err
}

func (s *productService) DeleteProduct(deleteProduct dto.DeleteProductRequest) error {
	product := entities.Product{
		ID: deleteProduct.ID,
	}

	return s.productRepo.DeleteProduct(product)
}
