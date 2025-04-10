package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", landingHandler)
	e.GET("/baseline", baselineHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
