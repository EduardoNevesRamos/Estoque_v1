package routes

import (
	controllers "github.com/DuduNeves/Estoque_v1/controller"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/", controllers.GetAllProducts)
			products.GET("/:id", controllers.GetProduct)
			products.POST("/", controllers.CreateProducts)
			products.PUT("/:id", controllers.UpdateProducts)
			products.DELETE("/:id", controllers.DeleteProducts)
		}
	}

	return router
}
