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
			products.GET("/", controllers.GetAllproducts)
		}
	}

	return router
}
