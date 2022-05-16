package controllers

import (
	"github.com/DuduNeves/Estoque_v1/database"
	"github.com/DuduNeves/Estoque_v1/models"
	"github.com/DuduNeves/Estoque_v1/service"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	db := database.GetDatabase()

	login := &models.Login{}
	err := ctx.ShouldBindJSON(login)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	user := &models.User{}
	dbError := db.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't find user",
		})
		return
	}

	if user.Password != service.SHA256Enconder(login.Password) {
		ctx.JSON(401, gin.H{
			"Error:": "Invalid credentials",
		})
		return
	}

	token, err := service.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"Error:": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}
