package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadProduct(g *gin.RouterGroup) {
	g.GET("/info", productInfo)
}

func productInfo(con *gin.Context) {
	con.JSON(http.StatusOK, gin.H{
		"result": "productInfo",
	})
}
