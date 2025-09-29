package http

import (
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"

	"net/http"
	"strconv"
)

func (s Server) GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if len(idStr) == 0 || err != nil {
		return echo.ErrBadRequest
	}

	user, err := s.service.user.GetUser(id)
	if err != nil {
		return err
	}

	response := user.ToResponse()

	c.JSON(http.StatusOK, response)
	return nil
}

func (s Server) GetUsers(c echo.Context) error {
	var params dto.GetUsersParams		

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	params.All = c.QueryParams()

	users, total, err := s.service.user.GetUsers(params)
	if err != nil {
		return err
	}

	var response dto.GetUsersResponse

	for _, user := range users {
		response.Data = append(response.Data, user.ToResponse())
	}

	response.Total = total;

	c.JSON(http.StatusOK, response)
	return nil
}

func (s Server) GetUserInfo(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims)
	id, err := strconv.Atoi(claims.Subject)

	if !ok || err != nil {
		return e.InternalServerError{Err: "error retrieving claims from context"}
	}

	user, err := s.service.user.GetUser(id)
	if err != nil {
		return err
	}

	response := user.ToResponse()

	return c.JSON(http.StatusOK, response)
}
