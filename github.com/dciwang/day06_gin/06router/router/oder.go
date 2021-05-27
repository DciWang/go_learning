package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadOrder(en *gin.RouterGroup) {
	en.GET("/info", orderInfo)
}

func orderInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "orderInfo",
	})
}
