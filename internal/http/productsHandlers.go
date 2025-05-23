package http

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/labstack/echo/v4"

	"net/http"
	"strconv"
	"strings"
)

func (s Server) CreateProduct(c echo.Context) error {
	var createProduct dto.CreateProductRequest

	if err := c.Bind(&createProduct); err != nil {
		return err
	}

	product, err := s.productService.CreateProduct(createProduct)
	if err != nil {
		return err
	}

	tags := strings.Split(product.Tags, ",")

	response := dto.CreateProductResponse{
		ID: product.ID,
		CreatedAt: product.CreatedAt,
		Name: product.Name,
		Description: product.Description,
		Vendor: product.Vendor,
		Price: product.Price,
		Tags: tags,
	}

	return c.JSON(http.StatusOK, response)
}

func (s Server) DeleteProduct(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if len(idStr) == 0 || err != nil { 
		return echo.ErrBadRequest
	}

	request := dto.DeleteProductRequest{
		ID: id,
	}

	if err := s.productService.DeleteProduct(request); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}
