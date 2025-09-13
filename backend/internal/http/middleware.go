package http

import (
	"github.com/dmxmss/e-commerce-app/entities"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
)

func GetAccessMiddleware(signingKey string, signingMethod string) echo.MiddlewareFunc {
	accessMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(signingKey),
		TokenLookup: "cookie:access_token",
		ContextKey: "user",
		SigningMethod: signingMethod,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entities.Claims)
		},
	})

	return accessMiddleware
}

func GetRefreshMiddleware(signingKey string, signingMethod string) echo.MiddlewareFunc {
	refreshMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(signingKey),
		TokenLookup: "cookie:refresh_token",
		ContextKey: "refresh",
		SigningMethod: signingMethod,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entities.Claims)
		},
	})

	return refreshMiddleware
}
