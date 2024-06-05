package handler

import (
	"go-backend-test/pkg/config"
	"go-backend-test/pkg/model"
	"go-backend-test/pkg/service"
	"go-backend-test/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	logger  *logrus.Logger
}

func NewUserHandler(service service.IUserService, maker token.Maker,
	config *config.Config, logger *logrus.Logger) IUserHandler {
	return &userHandler{
		service: service,
		maker:   maker,
		config:  config,
		logger:  logger,
	}
}

// Login implements IUserHandler.
func (u *userHandler) Login(ctx *gin.Context) {
	var loginRequest model.LoginRequest

	// JSON verisini bağlama (bind) işlemi ve hata loglaması
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":     err.Error(),
			"method":    ctx.Request.Method,
			"client_ip": "78.163.98.42",
		}).Error("Failed to bind JSON in Login")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kullanıcı doğrulama ve hata loglaması
	user, err := u.service.LoginUserService(loginRequest.UsernameOrMail, loginRequest.Password)
	if err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":    err.Error(),
			"username": loginRequest.UsernameOrMail,
			"method":   ctx.Request.Method,
		}).Error("Failed to authenticate user in Login")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Giriş başarılı loglama
	u.logger.WithFields(logrus.Fields{
		"userID":    user.ID,
		"method":    ctx.Request.Method,
		"client_ip": ctx.ClientIP(),
	}).Info("Login successful")

	// Token oluşturma ve hata loglaması
	token, err := u.maker.CreateToken(user.ID, user.Name, u.config.AccessTokenDuration)
	if err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":     err.Error(),
			"userID":    user.ID,
			"method":    ctx.Request.Method,
			"client_ip": ctx.ClientIP(),
		}).Error("Failed to create token in Login")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Token oluşturma başarılı loglama
	u.logger.WithFields(logrus.Fields{
		"userID": user.ID,
		"token":  token,
	}).Info("Token created successfully in Login")

	// Başarılı yanıt
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Register implements IUserHandler.
func (u *userHandler) Register(ctx *gin.Context) {
	var registerRequest model.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":     err.Error(),
			"method":    ctx.Request.Method,
			"client_ip": "78.163.98.42",
		}).Error("Failed to bind JSON in Register")
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
		u.logger.WithFields(logrus.Fields{
			"error":    err.Error(),
			"email":    registerRequest.Email,
			"username": registerRequest.Username,
		}).Error("Failed to register user")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.logger.WithFields(logrus.Fields{
		"email":    registerRequest.Email,
		"username": registerRequest.Username,
	}).Info("User registered successfully")
	ctx.JSON(http.StatusOK, gin.H{"success": "user created successfully"})
}

// DeleteUserHandler implements IUserHandler.
func (u *userHandler) DeleteUserHandler(ctx *gin.Context) {
	payload, err := u.maker.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":     err.Error(),
			"method":    ctx.Request.Method,
			"client_ip": ctx.ClientIP(),
		}).Error("Unauthorized access attempt in DeleteUserHandler")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	_, err = u.service.GetUserService(payload.ID)
	if err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":  err.Error(),
			"userID": payload.ID,
		}).Error("Failed to get user in DeleteUserHandler")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.service.DeleteService(payload.ID); err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":  err.Error(),
			"userID": payload.ID,
		}).Error("Failed to delete user")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	u.logger.WithField("userID", payload.ID).Info("User deleted successfully")
	ctx.JSON(http.StatusOK, gin.H{"success": "user deleted successfully"})
}

// GetUserHandler implements IUserHandler.
func (u *userHandler) GetUserHandler(ctx *gin.Context) {
	payload, err := u.maker.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":     err.Error(),
			"method":    ctx.Request.Method,
			"client_ip": "78.163.98.42",
		}).Error("Unauthorized access attempt in GetUserHandler")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := u.service.GetUserService(payload.ID)
	if err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":  err.Error(),
			"userID": payload.ID,
		}).Error("Failed to get user in GetUserHandler")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.logger.WithField("userID", payload.ID).Info("User fetched successfully")
	ctx.JSON(http.StatusOK, gin.H{"success": user})
}

// UpdateUserHandler implements IUserHandler.
func (u *userHandler) UpdateUserHandler(ctx *gin.Context) {
	payload, err := u.maker.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":     err.Error(),
			"method":    ctx.Request.Method,
			"client_ip": "78.163.98.42",
		}).Error("Unauthorized access attempt in UpdateUserHandler")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	_, err = u.service.GetUserService(payload.ID)
	if err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":  err.Error(),
			"userID": payload.ID,
		}).Error("Failed to get user in UpdateUserHandler")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := ctx.Bind(&user); err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":  err.Error(),
			"userID": payload.ID,
		}).Error("Failed to bind user data in UpdateUserHandler")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := u.service.UpdateService(user); err != nil {
		u.logger.WithFields(logrus.Fields{
			"error":  err.Error(),
			"userID": payload.ID,
		}).Error("Failed to update user")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	u.logger.WithField("userID", payload.ID).Info("User updated successfully")
	ctx.JSON(http.StatusOK, gin.H{"success": "user updated successfully"})
}
