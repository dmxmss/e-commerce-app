package http

import (
	e "github.com/dmxmss/e-commerce-app/error"
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
)

func (s Server) ErrorHandler(err error, c echo.Context) {
	var code int
	var msg string

	switch err.(type) {
	case e.UserAlreadyExists:
		code = http.StatusConflict
		msg = err.Error()
	case e.InvalidUserId:
		code = http.StatusBadRequest
		msg = err.Error()
	case e.InvalidCredentials:
		code = http.StatusUnauthorized
		msg = err.Error()
	case e.DbRecordNotFound:
		code = http.StatusNotFound
		msg = err.Error()
	case e.InvalidInputError:
		code = http.StatusBadRequest
		msg = err.Error()
	case *echo.HTTPError:
		code = err.(*echo.HTTPError).Code
		msg = fmt.Sprintf("%v", err.(*echo.HTTPError).Message)
	default:
		code = http.StatusInternalServerError
		msg = "internal server error"
	}

	if !c.Response().Committed {
		c.JSON(code, msg)
	}
}
