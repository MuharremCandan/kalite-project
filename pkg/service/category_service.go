package service

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/repository"

	"github.com/google/uuid"
)

type ICategoryService interface {
	CreateCategoryService(category *model.Category) error
	DeleteCategoryService(categoryID uuid.UUID) error
	UpdateCategoryService(category *model.Category) error
	GetCategoryService(categoryID uuid.UUID) (model.Category, error)
	GetCategoriesService() ([]model.Category, error)
}

type categoryService struct {
	repo repository.ICategoryRepository
}

func NewCategoryService(repo repository.ICategoryRepository) ICategoryService {
	return &categoryService{repo: repo}
}

// CreateCategoryService implements ICategoryService.
func (c *categoryService) CreateCategoryService(category *model.Category) error {
	return c.repo.Create(category)
}

// DeleteCategoryService implements ICategoryService.
func (c *categoryService) DeleteCategoryService(categoryID uuid.UUID) error {
	return c.repo.Delete(categoryID)
}

// GetCategoriesService implements ICategoryService.
func (c *categoryService) GetCategoriesService() ([]model.Category, error) {
	return c.repo.GetCategories()
}

// GetCategoryService implements ICategoryService.
func (c *categoryService) GetCategoryService(categoryID uuid.UUID) (model.Category, error) {
	return c.repo.GetCategory(categoryID)
}

// UpdateCategoryService implements ICategoryService.
func (c *categoryService) UpdateCategoryService(category *model.Category) error {
	return c.repo.Update(category)
}
