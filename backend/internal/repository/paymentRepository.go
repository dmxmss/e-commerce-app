package repository

import (
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/stripe/stripe-go/v82"

	"context"
	"fmt"
)

type PaymentRepository interface {
	CreatePayment(context.Context, int64, string, map[string]any) (*entities.Payment, error)
	GetPayment(context.Context, string, map[string]any) (*entities.Payment, error)
	CancelPayment(context.Context, string) error
}

type stripeRepository struct {
	sc *stripe.Client
}

func NewPaymentRepository(key string) PaymentRepository {
	sc := stripe.NewClient(key)

	return &stripeRepository{sc}
}

func (r *stripeRepository) CreatePayment(ctx context.Context, amount int64, currency string, metadata map[string]any) (*entities.Payment, error) {
	params := &stripe.PaymentIntentCreateParams{
		Amount: stripe.Int64(amount),
		Currency: stripe.String(string(currency)),
		AutomaticPaymentMethods: &stripe.PaymentIntentCreateAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	payment, err := r.sc.V1PaymentIntents.Create(ctx, params)	
	if err != nil {
		return nil, e.PaymentFailed{Err: fmt.Sprintf("error creating payment: %s", err)}
	}

	return &entities.Payment{
		ID: payment.ID,
		Currency: string(payment.Currency),
		AmountPaid: payment.Amount,
		Metadata: map[string]any{
			"clientSecret": payment.ClientSecret,
		},
	}, nil
}

func (r *stripeRepository) GetPayment(ctx context.Context, id string, metadata map[string]any) (*entities.Payment, error) {
	params := &stripe.PaymentIntentRetrieveParams{} 
	payment, err := r.sc.V1PaymentIntents.Retrieve(ctx, id, params)
	if err != nil {
		return nil, e.PaymentFailed{Err: fmt.Sprintf("error retrieving payment: %s", err)}
	}

	return &entities.Payment{
		ID: payment.ID,
		Currency: string(payment.Currency),
		AmountPaid: payment.Amount,
		Metadata: map[string]any{
			"clientSecret": payment.ClientSecret,
		},
	}, nil
}

func (r *stripeRepository) CancelPayment(ctx context.Context, id string) error {
	params := &stripe.PaymentIntentCancelParams{}
	_, err := r.sc.V1PaymentIntents.Cancel(ctx, id, params)
	if err != nil {
		return e.InternalServerError{Err: fmt.Sprintf("error cancelling payment: %s", err)}
	}

	return nil
}
