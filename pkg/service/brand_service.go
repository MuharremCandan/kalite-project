package service

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/repository"

	"github.com/google/uuid"
)

type IBrandService interface {
	CreateBrandService(brand model.Brand) error
	DeleteBrandService(brandID uuid.UUID) error
	UpdateBrandService(brand model.Brand) error
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
func (b *brandService) CreateBrandService(brand model.Brand) error {
	return b.repo.Create(brand)
}

// DeleteBrandService implements IBrandService.
func (b *brandService) DeleteBrandService(brandID uuid.UUID) error {
	return b.repo.Delete(brandID)
}

// GetBrandService implements IBrandService.
func (b *brandService) GetBrandService(brandID uuid.UUID) (model.Brand, error) {
	return b.repo.GetBrand(brandID)
}

// GetBrandsService implements IBrandService.
func (b *brandService) GetBrandsService() ([]model.Brand, error) {
	return b.repo.GetBrands()
}

// UpdateBrandService implements IBrandService.
func (b *brandService) UpdateBrandService(brand model.Brand) error {
	return b.repo.Update(brand)
}
