package handler

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	logger  *logrus.Logger
}

func NewProductHandler(service service.IProductService, logger *logrus.Logger) IProductHandler {
	return &productHandler{
		service: service,
		logger:  logger,
	}
}

// GetProductsByBrand implements IProductHandler.
func (p *productHandler) GetProductsByBrand(ctx *gin.Context) {
	var brandIDStr string
	if err := ctx.ShouldBindJSON(&brandIDStr); err != nil {
		p.logger.WithError(err).Error("Failed to bind JSON in GetProductsByBrand")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	brandID, err := uuid.Parse(brandIDStr)
	if err != nil {
		p.logger.WithError(err).Error("Failed to parse brandID in GetProductsByBrand")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	products, err := p.service.GetProductsByBrandService(brandID)
	if err != nil {
		p.logger.WithError(err).Error("Failed to get products by brand in GetProductsByBrand")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	p.logger.WithField("brandID", brandID).Info("Successfully retrieved products by brand")
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductsByCategory implements IProductHandler.
func (p *productHandler) GetProductsByCategory(ctx *gin.Context) {
	var categoryIDStr string
	if err := ctx.ShouldBindJSON(&categoryIDStr); err != nil {
		p.logger.WithError(err).Error("Failed to bind JSON in GetProductsByCategory")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		p.logger.WithError(err).Error("Failed to parse categoryID in GetProductsByCategory")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	products, err := p.service.GetProductsByCategoryService(categoryID)
	if err != nil {
		p.logger.WithError(err).Error("Failed to get products by category in GetProductsByCategory")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	p.logger.WithField("categoryID", categoryID).Info("Successfully retrieved products by category")
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductsByCategoryAndBrand implements IProductHandler.
func (p *productHandler) GetProductsByCategoryAndBrand(ctx *gin.Context) {
	var categoryID uuid.UUID
	var brandID uuid.UUID
	if err := ctx.ShouldBindJSON(&categoryID); err != nil {
		p.logger.WithError(err).Error("Failed to bind JSON for categoryID in GetProductsByCategoryAndBrand")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&brandID); err != nil {
		p.logger.WithError(err).Error("Failed to bind JSON for brandID in GetProductsByCategoryAndBrand")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	products, err := p.service.GetProductsByCategoryAndBrandService(categoryID, brandID)
	if err != nil {
		p.logger.WithError(err).Error("Failed to get products by category and brand in GetProductsByCategoryAndBrand")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	p.logger.WithFields(logrus.Fields{
		"categoryID": categoryID,
		"brandID":    brandID,
	}).Info("Successfully retrieved products by category and brand")
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

// CreateProduct implements IProductHandler.
func (p *productHandler) CreateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		p.logger.WithError(err).Error("Failed to bind JSON in CreateProduct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.service.CreateProductService(&product); err != nil {
		p.logger.WithError(err).Error("Failed to create product in CreateProduct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.logger.WithField("productID", product.ID).Info("Successfully created product")
	ctx.JSON(http.StatusOK, gin.H{"success": "product created successfully"})
}

// GetProduct implements IProductHandler.
func (p *productHandler) GetProduct(ctx *gin.Context) {
	productIDStr := ctx.Params.ByName("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		p.logger.WithError(err).Error("Failed to parse productID in GetProduct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := p.service.GetProductService(productID)
	if err != nil {
		p.logger.WithError(err).Error("Failed to get product in GetProduct")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	p.logger.WithField("productID", productID).Info("Successfully retrieved product")
	ctx.JSON(http.StatusOK, gin.H{"success": product})
}

// DeleteProduct implements IProductHandler.
func (p *productHandler) DeleteProduct(ctx *gin.Context) {
	productIDStr := ctx.Params.ByName("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		p.logger.WithError(err).Error("Failed to parse productID in DeleteProduct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.service.DeleteProductService(productID); err != nil {
		p.logger.WithError(err).Error("Failed to delete product in DeleteProduct")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	p.logger.WithField("productID", productID).Info("Successfully deleted product")
	ctx.JSON(http.StatusOK, gin.H{"success": "product deleted successfully"})
}

// GetProducts implements IProductHandler.
func (p *productHandler) GetProducts(ctx *gin.Context) {
	products, err := p.service.GetProductsService()
	if err != nil {
		p.logger.WithError(err).Error("Failed to get products in GetProducts")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	p.logger.Info("Successfully retrieved products")
	ctx.JSON(http.StatusOK, gin.H{"success": products})
}

// UpdateProduct implements IProductHandler.
func (p *productHandler) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		p.logger.WithError(err).Error("Failed to bind JSON in UpdateProduct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productIDStr := ctx.Params.ByName("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		p.logger.WithError(err).Error("Failed to parse productID in UpdateProduct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.service.UpdateProductService(productID, product); err != nil {
		p.logger.WithError(err).Error("Failed to update product in UpdateProduct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.logger.WithField("productID", productID).Info("Successfully updated product")
	ctx.JSON(http.StatusOK, gin.H{"success": "product updated successfully"})
}
