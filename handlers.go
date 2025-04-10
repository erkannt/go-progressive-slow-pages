package main

import (
	"time"

	"github.com/labstack/echo/v4"
)

func landingHandler(c echo.Context) error {
	return page(landing()).Render(c.Request().Context(), c.Response().Writer)
}

func baselineHandler(c echo.Context) error {
	results := getResults()
	return page(baseline(results)).Render(c.Request().Context(), c.Response().Writer)
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
