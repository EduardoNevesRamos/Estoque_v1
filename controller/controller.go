package controllers

import "github.com/gin-gonic/gin"

func GetAllproducts(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"Response:": "ok",
	})
}
