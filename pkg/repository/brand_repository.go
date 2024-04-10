package repository

import (
	"go-backend-test/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IBrandRepository interface {
	Create(brand *model.Brand) error
	Delete(brandID uuid.UUID) error
	Update(brand *model.Brand) error
	GetBrand(brandID uuid.UUID) (model.Brand, error)
	GetBrands() ([]model.Brand, error)
}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) IBrandRepository {
	return &brandRepository{db: db}
}

// Create implements IBrandRepository.
func (b *brandRepository) Create(brand *model.Brand) error {
	return b.db.Create(brand).Error
}

// Delete implements IBrandRepository.
func (b *brandRepository) Delete(brandID uuid.UUID) error {
	return b.db.Where("id = ? ", brandID).Delete(brandID).Error
}

// GetBrand implements IBrandRepository.
func (b *brandRepository) GetBrand(brandID uuid.UUID) (model.Brand, error) {
	var brand model.Brand
	err := b.db.Where("id =?", brandID).First(&brand).Error
	return brand, err
}

// GetBrands implements IBrandRepository.
func (b *brandRepository) GetBrands() ([]model.Brand, error) {
	var brands []model.Brand
	err := b.db.Find(&brands).Error
	return brands, err
}

// Update implements IBrandRepository.
func (b *brandRepository) Update(brand *model.Brand) error {
	return b.db.Save(&brand).Error
}
