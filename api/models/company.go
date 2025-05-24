package models

import (
	"time"

	"gorm.io/gorm"
)

// Company represents a company entity
type Company struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:255;not null" validate:"required"`
	Address   string         `json:"address" gorm:"size:500"`
	Phone     string         `json:"phone" gorm:"size:50"`
	Email     string         `json:"email" gorm:"size:255;not null;unique" validate:"required,email"`
	Website   string         `json:"website" gorm:"size:255"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete support
}

// TableName specifies the table name for the Company model
func (Company) TableName() string {
	return "companies"
}

// Companies is a slice of Company
type Companies []Company
