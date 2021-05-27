package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLFiles("./index.html")
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.html",nil)
	})
	engine.POST("/login", func(context *gin.Context) {
		type User struct {
			Username string `json:"username" form:"username"`
			Password string	`json:"password" form:"password"`
		}
		var u User
		if err := context.ShouldBind(&u); err != nil {
			context.JSON(http.StatusBadRequest,gin.H{
				"result":err.Error(),
			})
		}else {
		context.JSON(http.StatusOK,gin.H{
			"resulet":u,
		})

		}
	})

	engine.Run(":9090")
}
