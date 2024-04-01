package handler

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IProductHandler interface {
	GetProduct(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
}

// yorum
type productHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) IProductHandler {
	return &productHandler{
		service: service,
	}
}

// CreateProduct implements IProductHandler.
func (p *productHandler) CreateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.BrandID = uuid.New()
	product.CategoryID = uuid.New()
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
