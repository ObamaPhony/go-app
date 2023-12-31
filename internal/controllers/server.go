package controllers

import (
	"github.com/gin-gonic/gin"
    "github.com/obamaphony/go-app/internal/models"
)

func handleRoot(c *gin.Context) {
	c.Status(404)
}

func Server(bindInterface string) {
	r := gin.Default()
	r.GET("/", handleRoot)
}
