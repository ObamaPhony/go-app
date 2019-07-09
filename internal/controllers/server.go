package controllers

import (
	"github.com/gin-gonic/gin"
)

func Server(bindInter string) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"payload": "PONG",
		})
	})

	r.Run(bindInter)
}
