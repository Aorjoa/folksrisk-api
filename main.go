package main

import (
	"net/http"

	"github.com/aorjoa/folksrisk-api/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Database
	db := service.Database()
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello")
	})
	e.Logger.Fatal(e.StartTLS(":8443", "certs/cert.pem", "certs/key.pem"))
}
