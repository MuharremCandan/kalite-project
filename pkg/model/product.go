package model

import "github.com/google/uuid"

// product represents product information
type Product struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `json:"category_id" gorm:"index"`
	BrandID     uuid.UUID `json:"brand_id" gorm:"index"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ImagesUrl   string    `json:"images"`
}
