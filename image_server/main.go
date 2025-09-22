package main

import (
	"github.com/labstack/echo/v4"

	"net/http"
)

func main() {
	e := echo.New()

	e.Static("/", "images")

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Image server is running")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
