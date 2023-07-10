package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"msg": "GET OPPENING",
			})
		})
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"POST": "GET OPPENING",
			})
		})
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"PUT": "GET OPPENING",
			})
		})
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"DELETE": "GET OPPENING",
			})
		})
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"GET": "GET OPPENINGS",
			})
		})
	}
}
