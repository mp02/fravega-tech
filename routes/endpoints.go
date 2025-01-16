package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mp02/fravega-tech/docs"
	"github.com/mp02/fravega-tech/interfaces"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(handler *interfaces.ProductHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", handler.CreateProduct)
			products.GET("", handler.GetProducts)
			products.DELETE("/:id", handler.DeleteProduct)
			products.PATCH("/:id", handler.UpdateProductByID)
			//accounts.GET("/:id/statement", GetStatement)
		}
	}
	router.GET("/api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // ../api/doc/index.html
	return router
}
