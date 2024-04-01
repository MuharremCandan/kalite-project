package handler

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IBrandHandler interface {
	GetBrandHandler(ctx *gin.Context)
	CreateBrandHandler(ctx *gin.Context)
	UpdateBrandHandler(ctx *gin.Context)
	DeleteBrandHandler(ctx *gin.Context)
	GetBrandsHandler(ctx *gin.Context)
}

type handler struct {
	service service.IBrandService
}

func NewBrandHandler(service service.IBrandService) IBrandHandler {
	return &handler{service: service}
}

// CreateBrandHandler implements IBrandHandler.
func (h *handler) CreateBrandHandler(ctx *gin.Context) {
	var brand model.Brand
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateBrandService(brand); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Brand Created"})
}

// DeleteBrandHandler implements IBrandHandler.
func (h *handler) DeleteBrandHandler(ctx *gin.Context) {
	brandIDStr := ctx.Params.ByName("id")
	brandID, err := uuid.Parse(brandIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = h.service.GetBrandService(brandID)
	if err != nil {
		if err.Error() == "brand not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.DeleteBrandService(brandID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// GetBrandHandler implements IBrandHandler.
func (h *handler) GetBrandHandler(ctx *gin.Context) {
	brandIDStr := ctx.Params.ByName("id")
	brandID, err := uuid.Parse(brandIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	brand, err := h.service.GetBrandService(brandID)
	if err != nil {
		if err.Error() == "brand not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": brand})
}

// GetBrandsHandler implements IBrandHandler.
func (h *handler) GetBrandsHandler(ctx *gin.Context) {
	brands, err := h.service.GetBrandsService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": brands})
}

// UpdateBrandHandler implements IBrandHandler.
func (h *handler) UpdateBrandHandler(ctx *gin.Context) {
	var brand model.Brand
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateBrandService(brand); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Brand Updated"})
}
