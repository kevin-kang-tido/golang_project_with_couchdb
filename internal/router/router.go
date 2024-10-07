package router

import (
	"github.com/gin-gonic/gin"
	"golang_with_couchdb2/internal/delivery"
)

func InitRoutes(productHandler *delivery.ProductHandler) *gin.Engine {
	router := gin.Default()

	// Grouping product routes
	productRouter := router.Group("/products")
	{
		productRouter.POST("", productHandler.CreateProduct)
		productRouter.GET("", productHandler.GetAllProducts)
		productRouter.GET("/:_id", productHandler.GetProductByID)
		productRouter.PUT("/:_id", productHandler.UpdateProductByID)
		productRouter.DELETE("/:_id", productHandler.DeleteProductByID)
	}

	return router
}
