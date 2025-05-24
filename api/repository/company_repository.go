package repository

import (
	"company-api/api/models"
	"errors"

	"gorm.io/gorm"
)

// CompanyRepository handles database operations for Company entity
type CompanyRepository struct {
	db *gorm.DB
}

// NewCompanyRepository creates a new repository instance
func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

// GetAll retrieves all companies from the database
func (r *CompanyRepository) GetAll() (models.Companies, error) {
	var companies models.Companies
	result := r.db.Find(&companies)
	return companies, result.Error
}

// GetByID retrieves a company by its ID
func (r *CompanyRepository) GetByID(id uint) (*models.Company, error) {
	var company models.Company
	result := r.db.First(&company, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil when no record found
		}
		return nil, result.Error
	}
	return &company, nil
}

// Create adds a new company to the database
func (r *CompanyRepository) Create(company *models.Company) error {
	result := r.db.Create(company)
	return result.Error
}

// Update modifies an existing company in the database
func (r *CompanyRepository) Update(company *models.Company) error {
	result := r.db.Save(company)
	return result.Error
}

// Delete removes a company from the database
func (r *CompanyRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Company{}, id)
	return result.Error
}
