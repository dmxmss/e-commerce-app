package http

import (
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
)

func (s Server) ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var msg string

	switch err.(type) {
	case e.UserAlreadyExists:
		code = http.StatusConflict
	case e.InvalidUserId:
		code = http.StatusBadRequest
	case e.InvalidCredentials:
		code = http.StatusUnauthorized
	case e.DbRecordNotFound:
		code = http.StatusNotFound
	case e.InvalidInputError:
		code = http.StatusBadRequest
	}

	msg = err.Error()

	if he, ok := err.(*echo.HTTPError); ok {
    code = he.Code
		msg = fmt.Sprintf("%v", he.Message)
	}

	if !c.Response().Committed {
		c.JSON(code, msg)
	}
}
