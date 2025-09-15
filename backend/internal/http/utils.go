package http

import (
	"net/http"
)

func setCookie(name string, path string, age int, value string) *http.Cookie {
	cookie := new(http.Cookie)

	cookie.Name = name
	cookie.Value = value
	cookie.Path = path
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.MaxAge = age

	return cookie
}
