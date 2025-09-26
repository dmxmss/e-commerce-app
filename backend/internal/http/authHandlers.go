package http

import (
	"github.com/dmxmss/e-commerce-app/entities"
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/dmxmss/e-commerce-app/internal/dto"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"

	"net/http"
)

func (s Server) SignUp(c echo.Context) error {
	var request dto.CreateUserRequest

	if err := c.Bind(&request); err != nil {
		return err
	}

	user, accessToken, refreshToken, err := s.service.auth.SignUp(request.Name, request.Email, request.Password)
	if err != nil {
		return err
	}

	accessCookie := setCookie("access_token", "/", s.conf.Auth.Access.Expiration, accessToken.Value)
	refreshCookie := setCookie("refresh_token", "/", s.conf.Auth.Refresh.Expiration, refreshToken.Value)

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	response := dto.SignupResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Admin: user.Admin,
	}

	return c.JSON(http.StatusOK, response)
}

func (s Server) LogIn(c echo.Context) error {
	var request dto.LoginRequest

	if err := c.Bind(&request); err != nil {
		return err
	}

	user, accessToken, refreshToken, err := s.service.auth.Login(request)
	if err != nil {
		return err
	}

	accessCookie := setCookie("access_token", "/", s.conf.Auth.Access.Expiration, accessToken.Value)
	refreshCookie := setCookie("refresh_token", "/", s.conf.Auth.Refresh.Expiration, refreshToken.Value)

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	response := dto.LoginResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Admin: user.Admin,
	}

	return c.JSON(http.StatusOK, response)
}

func (s Server) LogOut(c echo.Context) error {
	accessCookie := setCookie("access_token", "/", -1, "")
	refreshCookie := setCookie("refresh_token", "/", -1, "")

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusNoContent, nil)
}

func (s Server) RefreshTokens(c echo.Context) error {
	token := c.Get("refresh").(*jwt.Token)
	r, ok := token.Claims.(*entities.Claims)

	if !ok {
		return e.InternalServerError{Err: "error retrieving claims from context"}
	}

	accessToken, refreshToken, err := s.service.auth.GenerateTokens(r.Admin, r.Subject)
	if err != nil {
		return err
	}

	accessCookie := setCookie("access_token", "/", s.conf.Auth.Access.Expiration, accessToken.Value)
	refreshCookie := setCookie("refresh_token", "/", s.conf.Auth.Refresh.Expiration, refreshToken.Value)

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, nil)
}
