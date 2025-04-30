package http

import (
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/dmxmss/e-commerce-app/entities"
	"github.com/labstack/echo/v4"

	"net/http"
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

	access, refresh, err := s.authService.GenerateTokens(user.ID)
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

	access, refresh, err := s.authService.GenerateTokens(user.ID)
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
	v := c.Get("refresh")
	r, ok := v.(entities.Claims)
	if v == nil || !ok {
		return echo.ErrUnauthorized
	}

	var request dto.RefreshTokensRequest	

	if err := c.Bind(&request); err != nil {
		return err
	}

	access, refresh, err := s.authService.GenerateTokens(r.UserId)
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
	v := c.Get("user")
	userClaims, ok := v.(entities.Claims)

	if v == nil || !ok {
		return echo.ErrUnauthorized
	}

	user, err := s.userService.GetUserInfo(userClaims.UserId)		
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
