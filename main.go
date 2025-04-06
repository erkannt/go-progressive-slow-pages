package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return landing().Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/search", func(c echo.Context) error {
		query := c.QueryParam("q")
		return search(query).Render(c.Request().Context(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
