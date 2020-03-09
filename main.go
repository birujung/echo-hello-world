package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/internal", internal)

	// Use 1323 as default port
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func internal(c echo.Context) error {
	return c.String(http.StatusOK, "Greeting 👋 from internally routed app!")
}
