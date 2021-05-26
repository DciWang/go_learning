package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	e := gin.Default()
	//e.LoadHTMLFiles("./templates/user/user.html","./templates/another/another.html")
	e.LoadHTMLGlob("templates/**/*")
	e.GET("/user/user", func(c *gin.Context) {

		c.HTML(http.StatusOK,"user/user.html",gin.H{
			"name":"wdc",
		})
	})

	e.GET("/another/another", func(c *gin.Context) {

		c.HTML(http.StatusOK,"another/another.html",gin.H{
			"name":"q",
		})
	})
	e.Run(":9090")
}
