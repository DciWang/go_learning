package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	e :=gin.Default()

	e.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
		"name":"wangduocong",
		"age":26,
		"result":"hello world",
		})
	})

	e.GET("/another", func(c *gin.Context) {
		type person struct {
			Name string `json:"name"`
			Age int 	`json:"age"`
			Result string `json:"result"`
		}
		p :=person{
			"q",25,"hello 多聪",
		}
		c.JSON(http.StatusOK,p)
	})

	e.Run(":9090")
}
