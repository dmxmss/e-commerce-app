package service

import (
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/dmxmss/e-commerce-app/internal/repository"
	"github.com/dmxmss/e-commerce-app/internal/dto"

	"context"
	"fmt"
)

type PaymentService interface {
	CreatePayment(context.Context, []int, string, map[string]any) (*entities.Payment, error)
	GetPayment(context.Context, string, map[string]any) (*entities.Payment, error)
}

type stripeServiceRepo struct { // repositories payment service needs
	payment repository.PaymentRepository
	product repository.ProductRepository
}

type stripeService struct {
	repo stripeServiceRepo	
}

func NewPaymentService(paymentRepo repository.PaymentRepository, productRepo repository.ProductRepository) PaymentService {
	return &stripeService{
		stripeServiceRepo{
			paymentRepo,
			productRepo,
		},
	}
}

func (s *stripeService) CreatePayment(ctx context.Context, productIds []int, currency string, metadata map[string]any) (*entities.Payment, error) {
	products, err := s.repo.product.GetProducts(dto.GetProductParams{IDs: productIds})
	if err != nil {
		return nil, err
	}
	
	var amount int64
	
	for _, product := range products {
		if product.Remaining == 0 {
			return nil, e.PaymentFailed{Err: fmt.Sprintf("no products left with id %d", product.ID)}
		}

		amount += int64(product.Price)
	}

	payment, err := s.repo.payment.CreatePayment(ctx, amount, currency, metadata)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *stripeService) GetPayment(ctx context.Context, id string, metadata map[string]any) (*entities.Payment, error) {
	payment, err := s.repo.payment.GetPayment(ctx, id, metadata)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
