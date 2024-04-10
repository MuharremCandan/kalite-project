package handler

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ICategoryHandler interface {
	CreateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	GetCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
}

type categoryHandler struct {
	service service.ICategoryService
}

func NewCategoryHandler(service service.ICategoryService) ICategoryHandler {
	return &categoryHandler{
		service: service,
	}
}

// CreateCategory implements ICategoryHandler.
func (c *categoryHandler) CreateCategory(ctx *gin.Context) {
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateCategoryService(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "category created successfully"})
}

// DeleteCategory implements ICategoryHandler.
func (c *categoryHandler) DeleteCategory(ctx *gin.Context) {
	categoryIDStr := ctx.Params.ByName("id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.DeleteCategoryService(categoryID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "category deleted successfully"})
}

// GetCategories implements ICategoryHandler.
func (c *categoryHandler) GetCategories(ctx *gin.Context) {
	categories, err := c.service.GetCategoriesService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": categories})
}

// GetCategory implements ICategoryHandler.
func (c *categoryHandler) GetCategory(ctx *gin.Context) {
	categoryIDStr := ctx.Params.ByName("id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := c.service.GetCategoryService(categoryID)
	if err != nil {
		if err.Error() == "category not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"success": category})
}

// UpdateCategory implements ICategoryHandler.
func (c *categoryHandler) UpdateCategory(ctx *gin.Context) {
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.UpdateCategoryService(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "category updated successfully"})
}
