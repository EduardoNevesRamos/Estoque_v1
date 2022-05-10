package controllers

import (
	"net/http"
	"strconv"

	"github.com/DuduNeves/api-Games/database"
	"github.com/DuduNeves/api-Games/models"
	"github.com/gin-gonic/gin"
)

func ShowAllGames(c *gin.Context) {
	db := database.GetDatabase()

	var games []models.Game
	err := db.Find(&games).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "cannot list games:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, games)
}

func DeleteGames(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Game{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete game: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func ShowGame(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be interger",
		})

		return 
	}

	db := database.GetDatabase()

	var game models.Game
	err = db.First(&game, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find game: " + err.Error(),
		})

		return
	}

	c.JSON(200, game)

}

func CreateGames(c *gin.Context) {
	db := database.GetDatabase()

	var game models.Game

	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON" + err.Error(),
		})
		return
	}

	err = db.Create(&game).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create game" + err.Error(),
		})
		return
	}

	c.JSON(200, game)
}

func EditGame(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Game

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot updated Game: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}
