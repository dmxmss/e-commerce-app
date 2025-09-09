package http

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"

	"net/http"
	"strconv"
	"log"
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


func (s Server) GetProducts(c echo.Context) error {
	var params dto.GetProductParams	

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	params.All = c.QueryParams()

	products, err := s.service.product.GetProducts(params)
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
		})
	}

	c.JSON(http.StatusOK, map[string]any {
		"data": response,
		"total": len(products),
	})
	return nil
}

func (s Server) GetProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if len(idStr) == 0 || err != nil { 
		return echo.ErrBadRequest
	}

	product, err := s.service.product.GetProduct(id)
	if err != nil {
		return err
	}

	response := dto.Product{
		ID: product.ID,
		Description: product.Description,
		CreatedAt: product.CreatedAt, 
		UpdatedAt: product.UpdatedAt,
		Name: product.Name,
		Vendor: product.Vendor,
		Remaining: product.Remaining,
		Price: product.Price,
		Category: product.CategoryID,
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (s Server) UpdateProduct(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims)
	if !ok {
		return e.InternalServerError{Err: "error retrieving claims from context"}
	} 
	_, err := strconv.Atoi(claims.Subject) // TODO: allow access only to vendor 
	if err != nil {
		return e.InternalServerError{Err: "error id conversion"}
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if len(idStr) == 0 || err != nil { 
		return echo.ErrBadRequest
	}

	var request dto.UpdateProductRequest

	if err := c.Bind(&request); err != nil {
		log.Printf("%s", err)
		return echo.ErrBadRequest
	}

	product, err := s.service.product.UpdateProduct(id, request)
	if err != nil {
		return err
	}

	response := dto.UpdateProductResponse{
		Description: product.Description,
		Name: product.Name,
		Remaining: product.Remaining,
		Price: product.Price,
		CategoryID: product.CategoryID,
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (s Server) DeleteProduct(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims)
	if !ok {
		return e.InternalServerError{Err: "error retrieving claims from context"}
	} 
	_, err := strconv.Atoi(claims.Subject) // TODO: allow access to vendor and admins
	if err != nil {
		return e.InternalServerError{Err: "error id conversion"}
	}

	idStr := c.QueryParam("id")
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
