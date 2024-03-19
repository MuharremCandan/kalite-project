package repository

import (
	"go-backend-test/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	Create(category model.Category) error
	Delete(categoryID uuid.UUID) error
	Update(category model.Category) error
	GetCategory(categoryID uuid.UUID) (model.Category, error)
	GetCategories() ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &categoryRepository{db: db}
}

// Create implements ICategoryRepository.
func (c *categoryRepository) Create(category model.Category) error {
	return c.db.Create(category).Error
}

// Delete implements ICategoryRepository.
func (c *categoryRepository) Delete(categoryID uuid.UUID) error {
	return c.db.Where("id =? ", categoryID).Delete(categoryID).Error
}

// GetCategories implements ICategoryRepository.
func (c *categoryRepository) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	err := c.db.Find(&categories).Error
	return categories, err
}

// GetCategory implements ICategoryRepository.
func (c *categoryRepository) GetCategory(categoryID uuid.UUID) (model.Category, error) {
	var category model.Category
	err := c.db.Where("id =?", categoryID).First(&category).Error
	return category, err
}

// Update implements ICategoryRepository.
func (c *categoryRepository) Update(category model.Category) error {
	return c.db.Save(&category).Error
}
