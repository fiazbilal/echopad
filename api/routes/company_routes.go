package routes

import (
	"company-api/api/controllers"
	"company-api/pkg/repository"
	"company-api/pkg/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterCompanyRoutes registers all routes for company operations
func RegisterCompanyRoutes(e *echo.Echo, db *gorm.DB) {
	// Create repository, service, and controller instances
	companyRepo := repository.NewCompanyRepository(db)
	companyService := service.NewCompanyService(companyRepo)
	companyController := controllers.NewCompanyController(companyService)

	// Company routes
	e.GET("/api/companies", companyController.GetCompanies)
	e.GET("/api/companies/:id", companyController.GetCompany)
	e.POST("/api/companies", companyController.CreateCompany)
	e.PUT("/api/companies/:id", companyController.UpdateCompany)
	e.DELETE("/api/companies/:id", companyController.DeleteCompany)
}
