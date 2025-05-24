package service

import (
	"company-api/api/models"
	"company-api/api/repository"
	"errors"
	"fmt"
	"strconv"
)

// CompanyService handles business logic for the Company entity
type CompanyService struct {
	repo *repository.CompanyRepository
}

// NewCompanyService creates a new service instance
func NewCompanyService(repo *repository.CompanyRepository) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

// GetAll returns all companies
func (s *CompanyService) GetAll() (models.Companies, error) {
	return s.repo.GetAll()
}

// GetByID returns a company by ID
func (s *CompanyService) GetByID(id string) (*models.Company, error) {
	// Convert string ID to uint
	uintID, err := parseID(id)
	if err != nil {
		return nil, err
	}
	
	company, err := s.repo.GetByID(uintID)
	if err != nil {
		return nil, err
	}
	
	if company == nil {
		return nil, errors.New("company not found")
	}
	
	return company, nil
}

// Create adds a new company
func (s *CompanyService) Create(company *models.Company) error {
	return s.repo.Create(company)
}

// Update modifies an existing company
func (s *CompanyService) Update(id string, company *models.Company) error {
	// Convert string ID to uint
	uintID, err := parseID(id)
	if err != nil {
		return err
	}
	
	// Check if the company exists
	existingCompany, err := s.repo.GetByID(uintID)
	if err != nil {
		return err
	}
	
	if existingCompany == nil {
		return errors.New("company not found")
	}
	
	// Update company ID and preserve timestamps
	company.ID = uintID
	company.CreatedAt = existingCompany.CreatedAt
	
	return s.repo.Update(company)
}

// Delete removes a company
func (s *CompanyService) Delete(id string) error {
	// Convert string ID to uint
	uintID, err := parseID(id)
	if err != nil {
		return err
	}
	
	// Check if the company exists
	existingCompany, err := s.repo.GetByID(uintID)
	if err != nil {
		return err
	}
	
	if existingCompany == nil {
		return errors.New("company not found")
	}
	
	return s.repo.Delete(uintID)
}

// Helper function to parse string ID to uint
func parseID(id string) (uint, error) {
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format: %w", err)
	}
	return uint(idInt), nil
}
