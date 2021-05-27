package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	e := gin.Default()
	defer e.Run(":9090")
	e.LoadHTMLFiles("./index.html")
	e.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	e.POST("/uploadfile", func(c *gin.Context) {
		c.Request.URL.Path = "/upload"
		e.HandleContext(c)

	})

	e.POST("/upload", func(c *gin.Context) {
		if file, err := c.FormFile("pic"); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"result": err.Error(),
			})
		} else {
			c.SaveUploadedFile(file, "./"+file.Filename)
		}
	})
}
