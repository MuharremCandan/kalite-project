package handler

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IProductHandler interface {
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	GetProductsByBrand(ctx *gin.Context)
	GetProductsByCategory(ctx *gin.Context)
	GetProductsByCategoryAndBrand(ctx *gin.Context)
}

type productHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) IProductHandler {
	return &productHandler{
		service: service,
	}
}

// GetProductsByBrand implements IProductHandler.
func (p *productHandler) GetProductsByBrand(ctx *gin.Context) {
	var brandIDStr string
	if err := ctx.ShouldBindJSON(&brandIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	brandID, err := uuid.Parse(brandIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	products, err := p.service.GetProductsByBrandService(brandID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductsByCategory implements IProductHandler.
func (p *productHandler) GetProductsByCategory(ctx *gin.Context) {
	var categoryIDStr string
	if err := ctx.ShouldBindJSON(&categoryIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	products, err := p.service.GetProductsByCategoryService(categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductsByCategoryAndBrand implements IProductHandler.
func (p *productHandler) GetProductsByCategoryAndBrand(ctx *gin.Context) {
	var categoryID uuid.UUID
	var brandID uuid.UUID
	if err := ctx.ShouldBindJSON(&categoryID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&brandID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	products, err := p.service.GetProductsByCategoryAndBrandService(categoryID, brandID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

// CreateProduct implements IProductHandler.
func (p *productHandler) CreateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.service.CreateProductService(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "product created successfully"})
}

// GetProduct implements IProductHandler.
func (p *productHandler) GetProduct(ctx *gin.Context) {
	productIDStr := ctx.Params.ByName("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := p.service.GetProductService(productID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": product})
}

// DeleteProduct implements IProductHandler.
func (p *productHandler) DeleteProduct(ctx *gin.Context) {
	productIDStr := ctx.Params.ByName("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.service.DeleteProductService(productID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "product deleted successfully"})
}

// GetProducts implements IProductHandler.
func (p *productHandler) GetProducts(ctx *gin.Context) {
	products, err := p.service.GetProductsService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": products})
}

// UpdateProduct implements IProductHandler.
func (p *productHandler) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productIDStr := ctx.Params.ByName("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.service.UpdateProductService(productID, product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "product updated successfully"})
}
