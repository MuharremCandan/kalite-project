package repository

import (
	"go-backend-test/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Create(product *model.Product) error
	Delete(productID uuid.UUID) error
	Update(productID uuid.UUID, product model.Product) error
	GetProduct(productID uuid.UUID) (model.Product, error)
	GetProducts() ([]model.Product, error)
	GetProductsByCategory(categoryID uuid.UUID) ([]model.Product, error)
	GetProductsByBrand(brandID uuid.UUID) ([]model.Product, error)
	GetProductsByCategoryAndBrand(categoryID uuid.UUID, brandID uuid.UUID) ([]model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db: db}
}

// Create implements IProductRepository.
func (p *productRepository) Create(product *model.Product) error {
	return p.db.Create(product).Error
}

// Delete implements IProductRepository.
func (p *productRepository) Delete(productID uuid.UUID) error {
	return p.db.Where("id =? ", productID).Delete(&model.Product{}).Error
}

// GetProduct implements IProductRepository.
func (p *productRepository) GetProduct(productID uuid.UUID) (model.Product, error) {
	var product model.Product
	err := p.db.Where("id =?", productID).First(&product).Error
	return product, err
}

// GetProducts implements IProductRepository.
func (p *productRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product
	err := p.db.Find(&products).Error
	return products, err
}

// GetProductsByBrand implements IProductRepository.
func (p *productRepository) GetProductsByBrand(brandID uuid.UUID) ([]model.Product, error) {
	var products []model.Product
	err := p.db.Where("brand_id =?", brandID).Find(&products).Error
	return products, err
}

// GetProductsByCategory implements IProductRepository.
func (p *productRepository) GetProductsByCategory(categoryID uuid.UUID) ([]model.Product, error) {
	var products []model.Product
	err := p.db.Where("category_id =?", categoryID).Find(&products).Error
	return products, err
}

// GetProductsByCategoryAndBrand implements IProductRepository.
func (p *productRepository) GetProductsByCategoryAndBrand(categoryID uuid.UUID, brandID uuid.UUID) ([]model.Product, error) {
	var products []model.Product
	err := p.db.Where("category_id =? and brand_id =?", categoryID, brandID).Find(&products).Error
	return products, err
}

// Update implements IProductRepository.
func (p *productRepository) Update(productID uuid.UUID, product model.Product) error {
	return p.db.Save(&product).Where("id =?", productID).Error
}
