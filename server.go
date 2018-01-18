package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true

	// Middlewares
	// - log everything
	e.Use(middleware.Logger())
	// - recover from panic errors
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hi.")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
