package handler

import (
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserHandler interface {
	CreateUserHandler(ctx *gin.Context)
	DeleteUserHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
	GetUserHandler(ctx *gin.Context)
}

type userHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) IUserHandler {
	return &userHandler{service: service}
}

// TODO: Swagger
// CreateUserHandler implements IUserHandler.
func (u *userHandler) CreateUserHandler(ctx *gin.Context) {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := u.service.CreateService(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": "user created successfully",
	})
}

// DeleteUserHandler implements IUserHandler.
func (u *userHandler) DeleteUserHandler(ctx *gin.Context) {

	userIDStr := ctx.Params.ByName("id")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = u.service.GetUserService(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.DeleteService(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": "user deleted successfully",
	})
}

// GetUserHandler implements IUserHandler.
func (u *userHandler) GetUserHandler(ctx *gin.Context) {

	userIDStr := ctx.Params.ByName("id")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := u.service.GetUserService(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": user,
	})
}

// UpdateUserHandler implements IUserHandler.
func (u *userHandler) UpdateUserHandler(ctx *gin.Context) {
	var user model.User

	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.UpdateService(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": "user updated successfully",
	})
}
