package http

import (
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"

	"net/http"
)

func (s Server) GetUserInfo(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims)

	if !ok {
		return e.InternalServerError{Err: "error retrieving claims from context"}
	}

	user, err := s.service.user.GetUserInfo(claims.Subject)
	if err != nil {
		return err
	}

	response := dto.GetUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Admin: user.Admin,
	}

	return c.JSON(http.StatusOK, response)
}
