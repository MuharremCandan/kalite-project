package router

import (
	"go-backend-test/pkg/api/handler"
	"go-backend-test/pkg/api/middleware"
	"go-backend-test/pkg/config"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type router struct {
	userHandler    handler.IUserHandler
	productHandler handler.IProductHandler
	brandHandler   handler.IBrandHandler
	cfg            *config.Config
}

func NewRouter(userHandler handler.IUserHandler,
	productHandler handler.IProductHandler,
	brandhandler handler.IBrandHandler, cfg *config.Config) *router {
	return &router{
		userHandler:    userHandler,
		productHandler: productHandler,
		brandHandler:   brandhandler,
		cfg:            cfg,
	}

}

func (h *router) SetUpRouter(r *gin.Engine) *gin.Engine {

	// Swagger UI için
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		publicRouter := v1.Group("/user")
		{
			publicRouter.POST("/login", h.userHandler.Login)
			publicRouter.POST("/register", h.userHandler.Register)
		}
		apiUser := v1.Group("/user")
		{
			// use authanticate to acces
			apiUser.Use(middleware.AuthMiddleware(h.cfg.SecretKey))
			apiUser.GET("/", h.userHandler.GetUserHandler)
			apiUser.PUT("/", h.userHandler.UpdateUserHandler)
			apiUser.DELETE("/", h.userHandler.DeleteUserHandler)
		}
		apiProduct := v1.Group("/product")
		{
			apiProduct.POST("/create", h.productHandler.CreateProduct)
			apiProduct.GET("/all", nil)
			apiProduct.GET("/:id", h.productHandler.GetProduct)
			apiProduct.PUT("/:id", nil)
			apiProduct.DELETE("/:id", nil)
		}
		apiCategory := v1.Group("/category")
		{
			apiCategory.POST("/create", nil)
			apiCategory.GET("/all", nil)
			apiCategory.GET("/:id", nil)
			apiCategory.PUT("/:id", nil)
			apiCategory.DELETE("/:id", nil)
		}
		apiBrand := v1.Group("/brand")
		{
			apiBrand.POST("/create", h.brandHandler.CreateBrandHandler)
			apiBrand.GET("/all", h.brandHandler.GetBrandsHandler)
			apiBrand.GET("/:id", h.brandHandler.GetBrandHandler)
			apiBrand.PUT("/:id", h.brandHandler.UpdateBrandHandler)
			apiBrand.DELETE("/:id", h.brandHandler.DeleteBrandHandler)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
