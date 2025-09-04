package http

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/labstack/echo/v4"

	"net/http"
)

func (s Server) CreatePayment(c echo.Context) error {
	var request dto.CreatePaymentRequest

	if err := c.Bind(&request); err != nil {
		return err	
	}

	ctx := c.Request().Context()

	payment, err := s.service.payment.CreatePayment(ctx, request.ProductIds, request.Currency, nil)
	if err != nil {
		return err
	}

	response := dto.CreatePaymentResponse{
		ClientSecret: payment.Metadata["clientSecret"].(string),
		PaymentId: payment.ID,
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (s Server) GetPayment(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return e.InvalidInputError{Err: "empty payment id"}
	}

	ctx := c.Request().Context()

	payment, err := s.service.payment.GetPayment(ctx, id, nil)
	if err != nil {
		return err
	}

	response := dto.CreatePaymentResponse{
		ClientSecret: payment.Metadata["clientSecret"].(string),
		PaymentId: payment.ID,
	}

	c.JSON(http.StatusOK, response)
	return nil
}
