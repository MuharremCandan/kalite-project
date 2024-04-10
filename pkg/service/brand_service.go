package service

import (
	"errors"
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IBrandService interface {
	CreateBrandService(brand *model.Brand) error
	DeleteBrandService(brandID uuid.UUID) error
	UpdateBrandService(brand *model.Brand) error
	GetBrandService(brandID uuid.UUID) (model.Brand, error)
	GetBrandsService() ([]model.Brand, error)
}

type brandService struct {
	repo repository.IBrandRepository
}

func NewBrandService(repo repository.IBrandRepository) IBrandService {
	return &brandService{repo: repo}
}

// CreateBrandService implements IBrandService.
func (b *brandService) CreateBrandService(brand *model.Brand) error {
	return b.repo.Create(brand)
}

// DeleteBrandService implements IBrandService.
func (b *brandService) DeleteBrandService(brandID uuid.UUID) error {
	return b.repo.Delete(brandID)
}

// GetBrandService implements IBrandService.
func (b *brandService) GetBrandService(brandID uuid.UUID) (model.Brand, error) {
	brand, err := b.repo.GetBrand(brandID)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return model.Brand{}, errors.New("brand not found")
		}
		return brand, err
	}
	return brand, nil
}

// GetBrandsService implements IBrandService.
func (b *brandService) GetBrandsService() ([]model.Brand, error) {
	return b.repo.GetBrands()
}

// UpdateBrandService implements IBrandService.
func (b *brandService) UpdateBrandService(brand *model.Brand) error {
	return b.repo.Update(brand)
}
