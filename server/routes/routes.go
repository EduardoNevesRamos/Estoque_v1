package routes

import (
	controllers "github.com/DuduNeves/Estoque_v1/controller"
	"github.com/DuduNeves/Estoque_v1/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		products := main.Group("products", middlewares.Auth())
		{
			products.GET("/", controllers.GetAllProducts)
			products.GET("/:id", controllers.GetProduct)
			products.POST("/", controllers.CreateProducts)
			products.PUT("/:id", controllers.UpdateProducts)
			products.DELETE("/:id", controllers.DeleteProducts)
		}

		user := main.Group("user")
		{
			user.GET("/", controllers.GetAllUsers)
			user.GET("/id", controllers.GetUser)
			user.POST("/", controllers.CreateUser)
		}

		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
	}

	return router
}
