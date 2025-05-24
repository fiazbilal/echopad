package controllers

import (
	"company-api/api/models"
	"company-api/pkg/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CompanyController handles company-related operations
type CompanyController struct {
	service *service.CompanyService
}

// NewCompanyController creates a new company controller
func NewCompanyController(service *service.CompanyService) *CompanyController {
	return &CompanyController{
		service: service,
	}
}

// GetCompanies returns all companies
func (c *CompanyController) GetCompanies(ctx echo.Context) error {
	companies, err := c.service.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to retrieve companies",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, companies)
}

// GetCompany returns a specific company by ID
func (c *CompanyController) GetCompany(ctx echo.Context) error {
	id := ctx.Param("id")
	
	company, err := c.service.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("Company with ID %s not found", id),
			"error":   err.Error(),
		})
	}
	
	return ctx.JSON(http.StatusOK, company)
}

// CreateCompany creates a new company
func (c *CompanyController) CreateCompany(ctx echo.Context) error {
	company := new(models.Company)
	
	if err := ctx.Bind(company); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}
	
	// Validate required fields
	if company.Name == "" || company.Email == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Name and Email are required fields",
		})
	}
	
	// Create the company
	err := c.service.Create(company)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create company",
			"error":   err.Error(),
		})
	}
	
	return ctx.JSON(http.StatusCreated, company)
}

// UpdateCompany updates an existing company
func (c *CompanyController) UpdateCompany(ctx echo.Context) error {
	id := ctx.Param("id")
	updatedCompany := new(models.Company)
	
	if err := ctx.Bind(updatedCompany); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}
	
	// Validate required fields
	if updatedCompany.Name == "" || updatedCompany.Email == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Name and Email are required fields",
		})
	}
	
	// Update the company
	err := c.service.Update(id, updatedCompany)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("Failed to update company with ID %s", id),
			"error":   err.Error(),
		})
	}
	
	return ctx.JSON(http.StatusOK, updatedCompany)
}

// DeleteCompany deletes a company
func (c *CompanyController) DeleteCompany(ctx echo.Context) error {
	id := ctx.Param("id")
	
	err := c.service.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("Failed to delete company with ID %s", id),
			"error":   err.Error(),
		})
	}
	
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Company with ID %s successfully deleted", id),
	})
}
