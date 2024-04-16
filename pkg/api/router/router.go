package router

import (
	"go-backend-test/pkg/api/handler"
	"go-backend-test/pkg/api/middleware"
	"go-backend-test/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type router struct {
	userHandler     handler.IUserHandler
	productHandler  handler.IProductHandler
	brandHandler    handler.IBrandHandler
	categoryHandler handler.ICategoryHandler
	cfg             *config.Config
}

func NewRouter(userHandler handler.IUserHandler,
	productHandler handler.IProductHandler,
	brandhandler handler.IBrandHandler,
	categoryHandler handler.ICategoryHandler,
	cfg *config.Config) *router {
	return &router{
		userHandler:     userHandler,
		productHandler:  productHandler,
		brandHandler:    brandhandler,
		categoryHandler: categoryHandler,
		cfg:             cfg,
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
			apiProduct.GET("/all", h.productHandler.GetProducts)
			apiProduct.GET("/:id", h.productHandler.GetProduct)
			apiProduct.PUT("/:id", h.productHandler.GetProduct)
			apiProduct.DELETE("/:id", h.productHandler.DeleteProduct)
			apiProduct.GET("/brand/:id", h.productHandler.GetProductsByBrand)
			apiProduct.GET("/category/:id", h.productHandler.GetProductsByCategory)
			apiProduct.GET("/category/:id/brand/:id", h.productHandler.GetProductsByCategoryAndBrand)
		}
		apiCategory := v1.Group("/category")
		{
			apiCategory.POST("/create", h.categoryHandler.CreateCategory)
			apiCategory.GET("/all", h.categoryHandler.GetCategories)
			apiCategory.GET("/:id", h.categoryHandler.GetCategory)
			apiCategory.PUT("/:id", h.categoryHandler.UpdateCategory)
			apiCategory.DELETE("/:id", h.categoryHandler.DeleteCategory)
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
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Check result is OK!",
		})
	})
	return r
}
