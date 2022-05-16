package controllers

import (
	"strconv"

	"github.com/DuduNeves/Estoque_v1/database"
	"github.com/DuduNeves/Estoque_v1/models"
	"github.com/DuduNeves/Estoque_v1/service"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	parseId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	user := models.User{}

	err = db.First(&user, parseId).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't find a user" + err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}

func GetAllUsers(ctx *gin.Context) {

	db := database.GetDatabase()
	user := []models.User{}

	err := db.Find(&user).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't find user by id: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, user)
}

func CreateUser(ctx *gin.Context) {
	db := database.GetDatabase()

	user := &models.User{}

	err := ctx.ShouldBindJSON(user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	user.Password = service.SHA256Enconder(user.Password)

	err = db.Create(user).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't create product" + err.Error(),
		})

		return
	}

	ctx.JSON(200, user)
}

func UpdateUser(ctx *gin.Context) {
	db := database.GetDatabase()

	user := &models.User{}

	err := ctx.ShouldBindJSON(user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	err = db.Save(user).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't update user" + err.Error(),
		})

		return
	}

	ctx.JSON(200, user)
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	parseId, err := strconv.Atoi(id)

	user := &models.User{}

	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	db := database.GetDatabase()

	err = db.Delete(user, parseId).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't update user" + err.Error(),
		})

		return
	}

	ctx.JSON(200, user)
}
