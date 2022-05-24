package routes

import (
	"github.com/DuduNeves/Estoque_v1/core/login"
	"github.com/DuduNeves/Estoque_v1/core/product"
	"github.com/DuduNeves/Estoque_v1/core/user"
	"github.com/DuduNeves/Estoque_v1/server/middlewares"
	"github.com/DuduNeves/Estoque_v1/server/routes/dependency"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	productsControllerWithDependencies := dependency.Products()

	{
		products := main.Group("products", middlewares.Auth())
		{
			products.GET("/", productsControllerWithDependencies.GetAllProducts)
			products.GET("/:id", productsControllerWithDependencies.GetProduct)
			products.POST("/", productsControllerWithDependencies.CreateProducts)
			products.PUT("/:id", productsControllerWithDependencies.UpdateProducts)
			products.DELETE("/:id", product.DeleteProducts)
		}

		users := main.Group("user")
		{
			users.GET("/", user.GetAllUsers)
			users.GET("/id", user.GetUser)
			users.POST("/", user.CreateUser)
			users.PUT("/:id", user.UpdateUser)
			users.DELETE("/:id", user.DeleteUser)
		}

		Login := main.Group("login")
		{
			Login.POST("/", login.Login)
		}
	}

	return router
}
