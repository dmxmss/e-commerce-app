package service

import (
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"gorm.io/gorm"

	"strings"
)

type ProductService interface {
	CreateProduct(string, string, string, int, []string) (*entities.Product, error)
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

func (s *productService) CreateProduct(name, description, vendor string, price int, tags []string) (*entities.Product, error) {
	concatTags := strings.Join(tags, ",")

	product := entities.Product{
		Name: name,
		Description: description,
		Vendor: vendor,
		Price: price,
		Tags: concatTags,
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
