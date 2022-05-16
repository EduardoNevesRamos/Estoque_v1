package controllers

import (
	"strconv"

	"github.com/DuduNeves/Estoque_v1/database"
	"github.com/DuduNeves/Estoque_v1/models"
	"github.com/gin-gonic/gin"
)

func GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	parseId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	product := models.Products{}

	err = db.First(&product, parseId).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't find a product" + err.Error(),
		})
		return
	}

	ctx.JSON(200, product)
}

func GetAllProducts(ctx *gin.Context) {

	db := database.GetDatabase()
	product := []models.Products{}

	err := db.Find(&product).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't find products by id: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, product)
}

func CreateProducts(ctx *gin.Context) {
	db := database.GetDatabase()

	product := &models.Products{}

	err := ctx.ShouldBindJSON(product)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	err = db.Create(product).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't create product" + err.Error(),
		})

		return
	}

	ctx.JSON(200, product)
}

func UpdateProducts(ctx *gin.Context) {
	db := database.GetDatabase()

	product := &models.Products{}

	err := ctx.ShouldBindJSON(product)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	err = db.Save(product).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't update product" + err.Error(),
		})

		return
	}

	ctx.JSON(200, product)
}

func DeleteProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	parseId, err := strconv.Atoi(id)

	product := &models.Products{}

	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	db := database.GetDatabase()

	err = db.Delete(product, parseId).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't update product" + err.Error(),
		})

		return
	}

	ctx.JSON(200, product)
}
