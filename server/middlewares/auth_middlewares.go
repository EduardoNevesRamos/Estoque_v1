package middlewares

import (
	"github.com/DuduNeves/Estoque_v1/service"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const Bearer_schena = "Bearer "
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(401)
		}

		token := header[len(Bearer_schena):]

		if !service.NewJWTService().ValidateToken(token) {
			ctx.AbortWithStatus(401)
		}
	}
}
