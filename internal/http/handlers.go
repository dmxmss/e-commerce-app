package http

import (
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"

	"net/http"
	"strconv"
)

func (s Server) SignUp(c echo.Context) error {
	var createUser dto.CreateUserRequest

	if err := c.Bind(&createUser); err != nil {
		return err
	}

	user, err := s.userService.CreateUser(createUser)
	if err != nil {
		return err
	}

	access, refresh, err := s.authService.GenerateTokens(strconv.Itoa(user.ID))
	if err != nil {
		return err
	}

	response := dto.CreateUserResponse{
		AccessToken: access,
		RefreshToken: refresh,
	}

	return c.JSON(http.StatusOK, response)
}

func (s Server) LogIn(c echo.Context) error {
	var loginUser dto.LoginUserRequest

	if err := c.Bind(&loginUser); err != nil {
		return err
	}

	user, err := s.userService.LogIn(loginUser)
	if err != nil {
		return err
	}

	access, refresh, err := s.authService.GenerateTokens(strconv.Itoa(user.ID))
	if err != nil {
		return err
	}

	response := dto.CreateUserResponse{
		AccessToken: access,
		RefreshToken: refresh,
	}

	return c.JSON(http.StatusOK, response)
}

func (s Server) RefreshTokens(c echo.Context) error {
	token := c.Get("refresh").(*jwt.Token)
	r, ok := token.Claims.(*entities.Claims)
	if !ok {
		return echo.ErrUnauthorized
	}

	var request dto.RefreshTokensRequest	

	if err := c.Bind(&request); err != nil {
		return err
	}

	access, refresh, err := s.authService.GenerateTokens(r.Subject)
	if err != nil {
		return err
	}

	response := dto.RefreshTokensResponse{
		AccessToken: access,
		RefreshToken: refresh,
	}

	return c.JSON(http.StatusOK, response)
}

func (s Server) GetUserInfo(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims, ok := token.Claims.(*entities.Claims)

	if !ok {
		return e.InternalServerError{Err: "error retrieving claims from context"}
	}

	user, err := s.userService.GetUserInfo(claims.Subject)		
	if err != nil {
		return err
	}

	response := dto.GetUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, response)
}
