package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return page(landing()).Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/search", func(c echo.Context) error {
		query := c.QueryParam("q")
		result_count := 3
		results := getResults(query)
		return page(search(query, result_count, results)).Render(c.Request().Context(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func getResults(query string) []string {
	var results []string
	if query != "" {
		results = []string{
			"a result",
			"another result",
			"and yet another result",
		}
	}
	return results
}
