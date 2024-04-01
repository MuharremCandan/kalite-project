package service

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/repository"

	"github.com/google/uuid"
)

type IProductService interface {
	CreateProductService(product *model.Product) error
	DeleteProductService(productID uuid.UUID) error
	UpdateProductService(product model.Product) error
	GetProductService(productID uuid.UUID) (model.Product, error)
	GetProductsService() ([]model.Product, error)
	GetProductsByCategoryService(categoryID uuid.UUID) ([]model.Product, error)
	GetProductsByBrandService(brandID uuid.UUID) ([]model.Product, error)
	GetProductsByCategoryAndBrandService(categoryID uuid.UUID, brandID uuid.UUID) ([]model.Product, error)
}

type productService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) IProductService {
	return &productService{repo: repo}
}

// CreateProductService implements IProductService.
func (p *productService) CreateProductService(product *model.Product) error {
	return p.repo.Create(product)
}

// DeleteProductService implements IProductService.
func (p *productService) DeleteProductService(productID uuid.UUID) error {
	return p.repo.Delete(productID)
}

// GetProductService implements IProductService.
func (p *productService) GetProductService(productID uuid.UUID) (model.Product, error) {
	return p.repo.GetProduct(productID)
}

// GetProductsByBrandService implements IProductService.
func (p *productService) GetProductsByBrandService(brandID uuid.UUID) ([]model.Product, error) {
	return p.repo.GetProductsByBrand(brandID)
}

// GetProductsByCategoryAndBrandService implements IProductService.
func (p *productService) GetProductsByCategoryAndBrandService(categoryID uuid.UUID, brandID uuid.UUID) ([]model.Product, error) {
	return p.repo.GetProductsByCategoryAndBrand(categoryID, brandID)
}

// GetProductsByCategoryService implements IProductService.
func (p *productService) GetProductsByCategoryService(categoryID uuid.UUID) ([]model.Product, error) {
	return p.repo.GetProductsByCategory(categoryID)
}

// GetProductsService implements IProductService.
func (p *productService) GetProductsService() ([]model.Product, error) {
	return p.repo.GetProducts()
}

// UpdateProductService implements IProductService.
func (p *productService) UpdateProductService(product model.Product) error {
	return p.repo.Update(product)
}
