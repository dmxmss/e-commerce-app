package http

import (
	"github.com/dmxmss/e-commerce-app/entities"

	"net/http"
)

func setCookieFromToken(name string, path string, age int, token entities.AuthToken) *http.Cookie {
	cookie := new(http.Cookie)

	cookie.Name = name
	cookie.Value = token.Value
	cookie.Path = path
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.MaxAge = age

	return cookie
}
