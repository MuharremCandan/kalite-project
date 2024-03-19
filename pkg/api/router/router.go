package router

import (
	"go-backend-test/pkg/api/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type router struct {
	userHandler handler.IUserHandler
}

func NewRouter() *router {
	return &router{}

}

func (h *router) SetUpRouter(r *gin.Engine) *gin.Engine {

	// Swagger UI için
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		apiUser := v1.Group("/user")
		{
			apiUser.POST("/create", h.userHandler.CreateUserHandler)
			//apiUser.GET("/all", h.userHandler.GetUserHandler)
			apiUser.GET("/:id", h.userHandler.GetUserHandler)
			apiUser.PUT("/:id", h.userHandler.UpdateUserHandler)
			apiUser.DELETE("/:id", h.userHandler.DeleteUserHandler)
		}
		apiProduct := v1.Group("/product")
		{
			apiProduct.POST("/create", nil)
			apiProduct.GET("/all", nil)
			apiProduct.GET("/:id", nil)
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
			apiBrand.POST("/create", nil)
			apiBrand.GET("/all", nil)
			apiBrand.GET("/:id", nil)
			apiBrand.PUT("/:id", nil)
			apiBrand.DELETE("/:id", nil)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
