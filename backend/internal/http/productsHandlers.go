package http

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"

	"net/http"
	"strconv"
)

func (s Server) CreateProduct(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims)
	if !ok {
		return e.InternalServerError{Err: "error retrieving claims from context"}
	} 
	vendorId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return e.InternalServerError{Err: "error id conversion"}
	}

	var request dto.CreateProductRequest

	if err := c.Bind(&request); err != nil {
		return err
	}

	product, err := s.service.product.CreateProduct(
		request.Name, 
		request.Description, 
		vendorId,
		request.Remaining,
		request.Price, 
		request.Category,
	)
	if err != nil {
		return err
	}

	response := dto.CreateProductResponse{
		ID: product.ID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Name: product.Name,
		Description: product.Description,
		Vendor: product.Vendor,
		Remaining: product.Remaining,
		Price: product.Price,
		Category: product.CategoryID,
	}

	return c.JSON(http.StatusOK, response)
}

func (s Server) GetUserProducts(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims)
	if !ok {
		return e.InternalServerError{Err: "error retrieving claims from context"}
	}

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return e.InternalServerError{Err: "error id conversion"}
	}

	products, err := s.service.product.GetUserProducts(id)
	if err != nil {
		return err
	}

	var response dto.GetProductsResponse

	for _, product := range products {
		response = append(response, dto.Product{
				ID: product.ID,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
				Name: product.Name,
				Description: product.Description,
				Vendor: product.Vendor,
				Remaining: product.Remaining,
				Price: product.Price,
				Category: product.CategoryID,
			},
		)
	}	

	return c.JSON(http.StatusOK, response)	
}

func (s Server) DeleteProduct(c echo.Context) error {
	idStr := c.QueryParam("id") // TODO: change
	id, err := strconv.Atoi(idStr)
	if len(idStr) == 0 || err != nil { 
		return echo.ErrBadRequest
	}

	if err := s.service.product.DeleteProduct(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}
