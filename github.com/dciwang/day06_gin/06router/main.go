package main

import (
	"github.com/gin-gonic/gin"
	"go_learning/github.com/dciwang/day06_gin/06router/router"
)

func main() {
	r := gin.Default()
	{
		orderGroup := r.Group("/order")
		router.LoadOrder(orderGroup)
	}
	{
		productGroup := r.Group("/product")
		router.LoadProduct(productGroup)
	}
	r.Run(":9090")
}
