package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	en := gin.Default()
	//en.LoadHTMLFiles("./index.html")
	en.LoadHTMLFiles("./index.html")
	en.GET("/query", func(context *gin.Context) {
		key := context.Query("key")
		/*	context.JSON(http.StatusOK,gin.H{
			"key":key,
		})*/
		context.HTML(http.StatusOK, "index.html", gin.H{
			"key": key,
		})
	})
	en.Run(":9090")
}
