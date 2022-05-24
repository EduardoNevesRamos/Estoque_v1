package login

import (
	"github.com/DuduNeves/Estoque_v1/database"
	"github.com/DuduNeves/Estoque_v1/database/entity"
	"github.com/DuduNeves/Estoque_v1/util/authorization"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	db := database.GetDatabase()

	login := &entity.Login{}
	err := ctx.ShouldBindJSON(login)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	user := &entity.User{}
	dbError := db.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't find user",
		})
		return
	}

	if user.Password != authorization.SHA256Enconder(login.Password) {
		ctx.JSON(401, gin.H{
			"Error:": "Invalid credentials",
		})
		return
	}

	token, err := authorization.NewJWTService().GenerateToken(user.ID)
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
