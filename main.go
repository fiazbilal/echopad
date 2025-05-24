package main

import (
	"company-api/api/routes"
	"company-api/config"
	"company-api/pkg/db"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	database, err := db.Connect(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run auto migrations
	err = db.Migrate(database)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome to Company API",
			"version": "1.0.0",
		})
	})

	// Register company routes
	routes.RegisterCompanyRoutes(e, database)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
