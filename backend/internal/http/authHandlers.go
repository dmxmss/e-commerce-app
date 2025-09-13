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

	_, accessToken, refreshToken, err := s.service.auth.SignUp(request.Name, request.Email, request.Password)
	if err != nil {
		return err
	}

	accessCookie := setCookieFromToken("access_token", "/", s.conf.Auth.Access.Expiration, *accessToken)
	refreshCookie := setCookieFromToken("refresh_token", "/", s.conf.Auth.Refresh.Expiration, *refreshToken)

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, nil)
}

func (s Server) LogIn(c echo.Context) error {
	var request dto.LoginRequest

	if err := c.Bind(&request); err != nil {
		return err
	}

	_, accessToken, refreshToken, err := s.service.auth.Login(request.Name, request.Password)
	if err != nil {
		return err
	}

	accessCookie := setCookieFromToken("access_token", "/", s.conf.Auth.Access.Expiration, *accessToken)
	refreshCookie := setCookieFromToken("refresh_token", "/", s.conf.Auth.Refresh.Expiration, *refreshToken)

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, nil)
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

	accessCookie := setCookieFromToken("access_token", "/", s.conf.Auth.Access.Expiration, *accessToken)
	refreshCookie := setCookieFromToken("refresh_token", "/", s.conf.Auth.Refresh.Expiration, *refreshToken)

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, nil)
}
