package routes

import (
	"github.com/DuduNeves/api-Games/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		games := main.Group("games")
		{
			games.GET("/:id", controllers.ShowGame)
			games.GET("/", controllers.ShowAllGames)
			games.POST("/", controllers.CreateGames)
			games.PUT("/", controllers.EditGame)
			games.DELETE("/:id", controllers.DeleteGames)

		}
	}

	return router
}
