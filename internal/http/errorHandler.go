package http

import (
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
)

func (s Server) ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "internal server error"

	switch err.(type) {
	case e.UserAlreadyExists:
		code = http.StatusConflict
		msg = err.Error()
	case e.UserNotFound:
		code = http.StatusNotFound
		msg = err.Error()
	case e.InvalidUserId:
		code = http.StatusBadRequest
		msg = err.Error()
	case e.InvalidCredentials:
		code = http.StatusUnauthorized
		msg = err.Error()
	}

	if he, ok := err.(*echo.HTTPError); ok {
    code = he.Code
		msg = fmt.Sprintf("%v", he.Message)
	}

	if !c.Response().Committed {
		c.JSON(code, msg)
	}
}
