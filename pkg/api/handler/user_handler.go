package handler

import (
	"go-backend-test/pkg/config"
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/service"
	"go-backend-test/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// yorum
type IUserHandler interface {
	DeleteUserHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
	GetUserHandler(ctx *gin.Context)
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userHandler struct {
	service service.IUserService
	maker   token.Maker
	config  *config.Config
}

func NewUserHandler(service service.IUserService, maker token.Maker, config *config.Config) IUserHandler {
	return &userHandler{
		service: service,
		maker:   maker,
		config:  config,
	}
}

// Login implements IUserHandler.
func (u *userHandler) Login(ctx *gin.Context) {
	var loginRequest model.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.service.LoginUserService(loginRequest.UsernameOrMail, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := u.maker.CreateToken(user.ID, user.Name, u.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Register implements IUserHandler.
func (u *userHandler) Register(ctx *gin.Context) {
	var registerRequest model.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.Wrap(err, "handler.User.Register"),
		})
		return
	}
	user := &model.User{
		Email:    registerRequest.Email,
		UserName: registerRequest.Username,
		Password: registerRequest.Password,
	}
	if err := u.service.RegisterUserService(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "user created successfully"})
}

// DeleteUserHandler implements IUserHandler.
func (u *userHandler) DeleteUserHandler(ctx *gin.Context) {
	payload, err := u.maker.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	_, err = u.service.GetUserService(payload.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := u.service.DeleteService(payload.ID); err != nil {
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
	payload, err := u.maker.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	user, err := u.service.GetUserService(payload.ID)
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
	payload, err := u.maker.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	_, err = u.service.GetUserService(payload.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
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
