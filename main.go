package main

import (
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return page(landing()).Render(c.Request().Context(), c.Response().Writer)
	})

	e.GET("/baseline", func(c echo.Context) error {
		results := getResults()
		return page(baseline(results)).Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func getResults() []string {
	results := []string{
		"a result",
		"another result",
		"and yet another result",
	}
	time.Sleep(2 * time.Second)
	return results
}
