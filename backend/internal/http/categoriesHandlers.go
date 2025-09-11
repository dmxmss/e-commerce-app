package http

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/labstack/echo/v4"

	"strconv"
	"net/http"
)

func (s Server) GetCategory(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if len(idStr) == 0 || err != nil {
		return echo.ErrBadRequest
	}

	category, err := s.service.category.GetCategory(id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
		"name": category.Name,
	})
	return nil
}

func (s Server) GetCategories(c echo.Context) error {
	var params dto.GetCategoriesParams

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	params.All = c.QueryParams()

	categories, err := s.service.category.GetCategories(params)
	if err != nil {
		return err
	}

	var response dto.GetCategoriesResponse
		
	for _, category := range categories {
		response = append(response, dto.Category{
			ID: category.ID,
			Name: category.Name,
		})
	}

	c.JSON(http.StatusOK, response)
	return nil
}
